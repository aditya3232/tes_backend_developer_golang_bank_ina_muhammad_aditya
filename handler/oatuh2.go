package handler

import (
	"context"
	"io"
	"net/http"

	oauth2_library "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/library/oauth2"
	oauth2_service "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/model/oauth2"
	"github.com/gin-gonic/gin"
)

type OAuth2Handler struct {
	oauth2Service oauth2_service.Service
}

func NewOAuth2Handler(oauth2Service oauth2_service.Service) *OAuth2Handler {
	return &OAuth2Handler{oauth2Service}
}

func (h *OAuth2Handler) GoogleLogin(c *gin.Context) {
	url := h.oauth2Service.GoogleLogin()
	// c.JSON(200, gin.H{"loginUrl": url})

	c.Redirect(http.StatusSeeOther, url)

}

func (h *OAuth2Handler) GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "randomstate" {
		c.String(http.StatusBadRequest, "States don't Match!!")
		return
	}

	code := c.Query("code")
	googleConfig := oauth2_library.GoogleConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Code-Token Exchange Failed")
		return
	}

	client := googleConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.String(http.StatusInternalServerError, "User Data Fetch Failed")
		return
	}
	defer resp.Body.Close()

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "JSON Parsing Failed")
		return
	}

	c.String(http.StatusOK, string(userData))
}