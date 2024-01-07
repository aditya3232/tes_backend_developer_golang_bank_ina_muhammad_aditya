package oauth2

import (
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/library/oauth2"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/model/users"
)

type Service interface {
	GoogleLogin(oauthStateString string) string
	GetUserInfo(email string) (bool, error)
}

type service struct {
	userRepository users.Repository
}

func NewService(userRepository users.Repository) *service {
	return &service{userRepository}
}

func (s *service) GoogleLogin(oauthStateString string) string {
	// state := "randomstate"
	url := oauth2.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthStateString)

	return url
}

func (s *service) GetUserInfo(email string) (bool, error) {
	// Check if the email exists in the database
	exists, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return false, err
	}
	return exists, nil
}
