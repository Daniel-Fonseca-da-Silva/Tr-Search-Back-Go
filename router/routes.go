package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	path := "/api/v1"
	v1 := router.Group(path)
	{
		v1.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET user",
			})
		})
		v1.POST("/user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST user",
			})
		})
		v1.DELETE("/user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE user",
			})
		})
		v1.PUT("/user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT user",
			})
		})
		v1.GET("/users", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET users",
			})
		})
	}
}
