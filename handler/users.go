package handler

import (
	"net/http"

	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/constant"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/helper"
	log_function "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/log"
	users_model "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/model/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// error asli ada di log, di response cuma menampilkan error umum untuk user

type UsersHandler struct {
	usersService users_model.Service
}

func NewUsersHandler(usersService users_model.Service) *UsersHandler {
	return &UsersHandler{usersService}
}

func (h *UsersHandler) GetAll(c *gin.Context) {
	users, err := h.usersService.GetAll()

	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusNotFound, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(users) == 0 {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := "Users not found"
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusNotFound, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, users_model.UserGetAllFormat(users))
	c.JSON(response.Meta.Code, response)

}

func (h *UsersHandler) GetByID(c *gin.Context) {
	var input users_model.UserGetOneByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	user, err := h.usersService.GetOne(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusNotFound, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	if user.ID == 0 {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := "Users not found"
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusNotFound, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, users_model.UserGetFormat(user))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) Create(c *gin.Context) {
	var input users_model.UserCreateInput

	err := c.ShouldBindWith(&input, binding.Form)

	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	newUser, err := h.usersService.Create(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessCreateData
	infoCode := http.StatusCreated
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusCreated, users_model.UserCreateFormat(newUser))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) Update(c *gin.Context) {
	var id users_model.UserGetOneByIdInput
	var input users_model.UserUpdateInput

	err := c.ShouldBindUri(&id)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	input.ID = id.ID

	err = c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	newUser, err := h.usersService.Update(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.FailedUpdateData
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessUpdateData
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, users_model.UserUpdateFormat(newUser))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) Delete(c *gin.Context) {
	var id users_model.UserGetOneByIdInput

	err := c.ShouldBindUri(&id)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	err = h.usersService.Delete(id)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.FailedDeleteData
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessDeleteData
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, nil)
	c.JSON(response.Meta.Code, response)
}
