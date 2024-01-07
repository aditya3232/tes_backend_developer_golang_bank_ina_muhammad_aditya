package oauth2

import "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/library/oauth2"

type Service interface {
	GoogleLogin(oauthStateString string) string
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) GoogleLogin(oauthStateString string) string {
	// state := "randomstate"
	url := oauth2.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthStateString)

	return url
}
