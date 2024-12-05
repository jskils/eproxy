package router

import (
	"eproxy/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		passwordQuery := c.Query("password")
		passwordHeader := c.GetHeader("password")
		password := os.Getenv("password")
		if passwordQuery != password && passwordHeader != password {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized: Invalid password",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	oapiGroup := r.Group("/oapi").Use(AuthMiddleware())

	{
		oapiGroup.POST("/run", api.Run)
		oapiGroup.POST("/change", api.Change)
		oapiGroup.POST("/release", api.Release)

	}
	return r
}
