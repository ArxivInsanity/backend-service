package auth

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func SetupAuth() {
	clientId := os.Getenv(OAUTH2_CLIENT_ID)
	secret := os.Getenv(OAUTH2_SECRET)
	goth.UseProviders(
		google.New(clientId, secret, os.Getenv(OAUTH2_REDIRECT_URL), "email", "profile"),
	)
}
