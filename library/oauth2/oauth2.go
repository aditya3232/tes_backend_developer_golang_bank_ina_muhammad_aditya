package oauth2

import (
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig oauth2.Config
}

var AppConfig Config

func GoogleConfig() oauth2.Config {
	AppConfig.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  config.CONFIG.GOOGLE_REDIRECT_URL,
		ClientID:     config.CONFIG.GOOGLE_CLIENT_ID,
		ClientSecret: config.CONFIG.GOOGLE_CLIENT_SECRET,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return AppConfig.GoogleLoginConfig
}
