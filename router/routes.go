package router

import (
	"github.com/Daniel-Fonseca-da-Silva/Tr-Search-Back-Go/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	path := "/api/v1"
	v1 := router.Group(path)
	{
		v1.GET("/user", handler.ShowUserHandler)

		v1.POST("/user", handler.CreateUserHandler)

		v1.DELETE("/user", handler.DeleteUserHandler)

		v1.PUT("/user", handler.UpdateUserHandler)

		v1.GET("/users", handler.ListUsersHandler)
	}
}
