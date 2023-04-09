package auth

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func SetupAuth(c *gin.Context) {
	clientId := os.Getenv(OAUTH2_CLIENT_ID)
	secret := os.Getenv(OAUTH2_SECRET)
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	redirectUrl := fmt.Sprint(scheme, "://", c.Request.Host, "/auth/google/callback")
	goth.UseProviders(
		google.New(clientId, secret, redirectUrl, "email", "profile"),
	)
}
