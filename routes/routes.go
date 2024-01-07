package routes

import (
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/connection"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/handler"
	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/model/users"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	usersRepository := users.NewRepository(connection.DatabaseMysql())

	// Initialize services
	usersService := users.NewService(usersRepository)

	// Initialize handlers
	usersHandler := handler.NewUsersHandler(usersService)

	// Configure routes
	api := router.Group("")

	usersRoutes := api.Group("")

	configureUsersRoutes(usersRoutes, usersHandler)

}

func configureUsersRoutes(group *gin.RouterGroup, handler *handler.UsersHandler) {
	group.GET("/users", handler.GetAll)
	group.POST("/users", handler.Create)
	group.GET("/users/:id", handler.GetByID)
	group.PUT("/users/:id", handler.Update)
	group.DELETE("/users/:id", handler.Delete)
}
