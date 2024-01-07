package routes

import (
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/connection"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/handler"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/model/tasks"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/model/users"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	usersRepository := users.NewRepository(connection.DatabaseMysql())
	taskRepository := tasks.NewRepository(connection.DatabaseMysql())

	// Initialize services
	usersService := users.NewService(usersRepository)
	taskService := tasks.NewService(taskRepository)

	// Initialize handlers
	usersHandler := handler.NewUsersHandler(usersService)
	taskHandler := handler.NewTasksHandler(taskService)

	// Configure routes
	api := router.Group("")

	usersRoutes := api.Group("")
	tasksRoutes := api.Group("")

	configureUsersRoutes(usersRoutes, usersHandler)
	configureTasksRoutes(tasksRoutes, taskHandler)

}

func configureUsersRoutes(group *gin.RouterGroup, handler *handler.UsersHandler) {
	group.GET("/users", handler.GetAll)
	group.POST("/users", handler.Create)
	group.GET("/users/:id", handler.GetByID)
	group.PUT("/users/:id", handler.Update)
	group.DELETE("/users/:id", handler.Delete)
}

func configureTasksRoutes(group *gin.RouterGroup, handler *handler.TasksHandler) {
	group.GET("/tasks", handler.GetAll)
	group.POST("/tasks", handler.Create)
	group.GET("/tasks/:id", handler.GetByID)
	group.PUT("/tasks/:id", handler.Update)
	group.DELETE("/tasks/:id", handler.Delete)
}
