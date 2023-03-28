package auth

import "github.com/gin-gonic/gin"

func AddAuthRoutes(router *gin.Engine) {
	router.GET("/auth/:provider", Redirect)
	router.GET("/auth/:provider/callback", HandleRedirect)
	router.Use(AuthenticationRequired())
	router.GET("/auth/isLoggedIn", IsLoggedIn)
	router.GET("/auth/getUserInfo", GetUserInfo)
}
