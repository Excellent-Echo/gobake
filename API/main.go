package main

import (
	"github.com/gin-gonic/gin"
	"go-bake/auth"
	"go-bake/config"
	"go-bake/handler"
	"go-bake/user"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.SetupConnection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func main() {
	r := gin.Default()

	r.GET("/users", handler.MiddleWare(userService, authService), userHandler.ShowUserHandler)
	r.GET("/users/:user_id", handler.MiddleWare(userService, authService), userHandler.GetUserByIDHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.PUT("/users/:user_id", handler.MiddleWare(userService, authService), userHandler.UpdateUserByIDHandler)
	r.DELETE("/users/:user_id", handler.MiddleWare(userService, authService), userHandler.DeleteUserByIDHandler)

	r.Run(":1204")

}
