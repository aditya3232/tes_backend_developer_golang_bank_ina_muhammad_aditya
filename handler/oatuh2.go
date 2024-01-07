package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/config"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/helper"
	oauth2_library "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/library/oauth2"
	log_function "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/log"
	oauth2_service "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/model/oauth2"
	"github.com/gin-gonic/gin"
)

type OAuth2Handler struct {
	oauth2Service oauth2_service.Service
}

func NewOAuth2Handler(oauth2Service oauth2_service.Service) *OAuth2Handler {
	return &OAuth2Handler{oauth2Service}
}

var (
	oauthStateString = config.CONFIG.OAUTH_STATE_STRING
	googleConfig     = oauth2_library.GoogleConfig()
	userInfo         = oauth2_service.UserInfo{}
)

func (h *OAuth2Handler) GoogleLogin(c *gin.Context) {
	url := h.oauth2Service.GoogleLogin(oauthStateString)
	c.JSON(http.StatusOK, gin.H{"login_url": url})
}

func (h *OAuth2Handler) GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		endpoint := c.Request.URL.Path
		message := "States don't Match!!"
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := "States don't Match!!"
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	code := c.Query("code")
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := "Code-Token Exchange Failed"
		errorCode := http.StatusInternalServerError
		ipAddress := c.ClientIP()
		errors := "Code-Token Exchange Failed"
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusInternalServerError, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	// Retrieve user information from https://www.googleapis.com/oauth2/v2/userinfo
	client := googleConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		endpoint := c.Request.URL.Path
		message := "User Data Fetch Failed"
		errorCode := http.StatusInternalServerError
		ipAddress := c.ClientIP()
		errors := "User Data Fetch Failed"
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusInternalServerError, nil)
		c.JSON(response.Meta.Code, response)
		return
	}
	defer resp.Body.Close()

	// Parse JSON data
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		endpoint := c.Request.URL.Path
		message := "JSON Parsing Failed"
		errorCode := http.StatusInternalServerError
		ipAddress := c.ClientIP()
		errors := "JSON Parsing Failed"
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusInternalServerError, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	// Check if the email exists in the database
	userInfoExists, err := h.oauth2Service.GetUserInfo(userInfo.Email)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := "Users Not Found"
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := "Users not found"
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusNotFound, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	if userInfoExists {
		userInfo.Expiry = token.Expiry.Format(time.RFC3339)
		userInfo.AccessTokenToken = token.AccessToken

		endpoint := c.Request.URL.Path
		message := "Login Success"
		infoCode := http.StatusOK
		ipAddress := c.ClientIP()
		log_function.Info(message, "", endpoint, infoCode, ipAddress)

		response := helper.APIResponse(message, http.StatusOK, oauth2_service.LoginFormat(userInfo))
		c.JSON(response.Meta.Code, response)

	} else {
		endpoint := c.Request.URL.Path
		message := "Email not authorized"
		errorCode := http.StatusUnauthorized
		ipAddress := c.ClientIP()
		errors := "Email not authorized"
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse("", http.StatusUnauthorized, nil)
		c.JSON(response.Meta.Code, response)
	}

}
