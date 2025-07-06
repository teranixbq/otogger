package config

import (
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func gapi() *oauth2.Config {
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("clientID"),
		ClientSecret: os.Getenv("clientSecret"),
		RedirectURL:  os.Getenv("redirectURL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/blogger",
		},
		Endpoint: google.Endpoint,
	}
	return oauthConfig
}
