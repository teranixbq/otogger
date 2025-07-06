package dto

import (
	"time"

	"golang.org/x/oauth2"
)

type AuthResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expired      time.Time `json:"expired"`
}

type RefreshResponse struct {
	AccessToken  string    `json:"access_token"`
	Expired      time.Time `json:"expired"`
}

func TokenResult(token oauth2.Token) AuthResponse {
	return AuthResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expired:      token.Expiry,
	}
}

func TokenRefreshResult(token oauth2.Token) RefreshResponse {
	return RefreshResponse{
		AccessToken:  token.AccessToken,
		Expired:      token.Expiry,
	}
}
