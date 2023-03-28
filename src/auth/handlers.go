package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/oov/gothic"
	"github.com/rs/zerolog/log"
)

type UserDetails struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	EmailId    string `json:"emailId"`
	ProfilePic string `json:"profilePic"`
}

type Claims struct {
	UserDetails
	jwt.RegisteredClaims
}

var JwtKey []byte = []byte(os.Getenv(JWT_SECRET))

// Auth godoc
// @Summary Endpoint for logging in the user using google Oauth 2
// @Schemes
// @Description Will redirect user to google OAuth consent screen
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /auth/google [get]
func Redirect(c *gin.Context) {
	err := gothic.BeginAuth(c.Param("provider"), c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

// Auth godoc
// @Summary Endpoint for handling the google OAuth callback
// @Schemes
// @Description Will handel the google OAuth call back and redirect to homepage
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /auth/google/callback [get]
func HandleRedirect(c *gin.Context) {
	user, err := gothic.CompleteAuth(c.Param("provider"), c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Debug().Msg("Retrieved user " + user.Email)
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserDetails: UserDetails{
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			EmailId:    user.Email,
			ProfilePic: user.AvatarURL,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.IndentedJSON(http.StatusInternalServerError, "Something went wrong when signing jwt token")
		log.Error().Msg("Error generating JWT token: " + err.Error())
		return
	}
	c.SetCookie(COOKIE_KEY, tokenString, 86400, "/", "", true, true)
	log.Debug().Msg("Redirecting to home page now ")
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/")
}

func extractClaimsFromJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Error().Msg("Invalid signature")
			return claims, errors.New(string(http.StatusUnauthorized))
		}
		log.Error().Msg("Bad request")
		return claims, errors.New(string(http.StatusBadRequest))
	}
	if !tkn.Valid {
		log.Error().Msg("Invalid token")
		return claims, errors.New(string(http.StatusUnauthorized))
	}
	return claims, nil

}

func IsValidLogin(c *gin.Context) bool {
	tokenStr, err := c.Cookie(COOKIE_KEY)
	if err != nil {
		log.Error().Msg("Failed to read from cookie")
		return false
	}
	_, err = extractClaimsFromJWT(tokenStr)
	if err != nil {
		log.Error().Msg("Failed to extract claim from cookie")
		return false
	}
	return true
}

// Auth godoc
// @Summary Endpoint for checking if user is logged in
// @Schemes
// @Description Checks if there is a cookie preset with the jwt token. If present, validates the token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 401 {object} string
// @Router /auth/isLoggedIn [get]
func IsLoggedIn(c *gin.Context) {
	if !IsValidLogin(c) {
		c.IndentedJSON(http.StatusUnauthorized, "User not logged in")
	} else {
		c.IndentedJSON(http.StatusOK, "User is logged in")
	}
}

// Auth godoc
// @Summary Endpoint for getting user details
// @Schemes
// @Description Checks if there is a cookie preset with the jwt token. If present, validates the token and then returns the user details
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} UserDetails
// @Failure 401 {object} string
// @Router /auth/getUserInfo [get]
func GetUserInfo(c *gin.Context) {
	tokenStr, err := c.Cookie(COOKIE_KEY)
	if err != nil {
		if err == http.ErrNoCookie {
			c.IndentedJSON(http.StatusUnauthorized, "User not logged in")
		} else {
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong when retrieving token")
		}
		return
	}
	claims, errorDetails := extractClaimsFromJWT(tokenStr)
	if errorDetails != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Something went wrong when getting user details")
		log.Debug().Msg("Error : " + err.Error())
	}
	c.IndentedJSON(http.StatusOK, claims.UserDetails)
}
