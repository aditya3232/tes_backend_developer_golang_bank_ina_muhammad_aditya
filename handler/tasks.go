package handler

import (
	"net/http"

	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/constant"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/helper"
	log_function "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/log"
	tasks_model "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/model/tasks"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type TasksHandler struct {
	tasksService tasks_model.Service
}

func NewTasksHandler(tasksService tasks_model.Service) *TasksHandler {
	return &TasksHandler{tasksService}
}

func (h *TasksHandler) GetAll(c *gin.Context) {
	tasks, err := h.tasksService.GetAll()

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

	if len(tasks) == 0 {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := "Tasks not found"
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

	response := helper.APIResponse(message, http.StatusOK, tasks_model.TaskGetAllFormat(tasks))
	c.JSON(response.Meta.Code, response)
}

func (h *TasksHandler) GetByID(c *gin.Context) {
	var input tasks_model.TaskGetOneByIdInput

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

	task, err := h.tasksService.GetOne(input)
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

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, tasks_model.TaskGetFormat(task))
	c.JSON(response.Meta.Code, response)
}

func (h *TasksHandler) Create(c *gin.Context) {
	var input tasks_model.TaskCreateInput

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

	newTask, err := h.tasksService.Create(input)
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

	response := helper.APIResponse(message, http.StatusCreated, tasks_model.TaskCreateFormat(newTask))
	c.JSON(response.Meta.Code, response)
}

func (h *TasksHandler) Update(c *gin.Context) {
	var id tasks_model.TaskGetOneByIdInput
	var input tasks_model.TaskUpdateInput

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

	newTask, err := h.tasksService.Update(input)
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

	response := helper.APIResponse(message, http.StatusOK, tasks_model.TaskUpdateFormat(newTask))
	c.JSON(response.Meta.Code, response)
}

func (h *TasksHandler) Delete(c *gin.Context) {
	var id tasks_model.TaskGetOneByIdInput

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

	err = h.tasksService.Delete(id)
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
