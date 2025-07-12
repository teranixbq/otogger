package service

import (
	"context"
	"errors"
	"gapi/auth/dto"

	"golang.org/x/oauth2"
)

type subService struct {
	oa *oauth2.Config
}

type ServiceInterface interface{
	RedirectUrl() (string,error)
	AccessToken(code string, ctx context.Context) (dto.AuthResponse,error)
	RefreshToken(RefreshToken string, ctx context.Context) (dto.RefreshResponse,error)
}

func NewAuthService(oa *oauth2.Config) ServiceInterface {
	return &subService{
		oa: oa,
	}
}

// after login with google u will redirect, and google will give u code in parameter
func (s *subService) RedirectUrl() (string, error) {
	url := s.oa.AuthCodeURL("state-token-123", oauth2.AccessTypeOffline)
	if url == "" {
		return "", errors.New("error : failed to generate redirect URL")
	}
	return url,nil
}

// get code from url redirect automaticly, and make sure to save refreshtoken
func (s *subService) AccessToken(code string, ctx context.Context) (dto.AuthResponse,error){
 	token,err := s.oa.Exchange(ctx,code)
	if err != nil {
		return dto.AuthResponse{},err
	}

	result := dto.TokenResult(*token)
	return result,nil
}

// using refreshtoken from first request oauth2.
func (s *subService) RefreshToken(RefreshToken string, ctx context.Context) (dto.RefreshResponse, error){
	tokenSource := s.oa.TokenSource(ctx,&oauth2.Token{RefreshToken:RefreshToken})

	newToken,err := tokenSource.Token()
	if err != nil {
		return dto.RefreshResponse{}, err
	}

	result := dto.TokenRefreshResult(*newToken)
	return  result,nil
}