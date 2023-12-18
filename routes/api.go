// Package routes register web routes
package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes register web routes
func RegisterWebRoutes(router *gin.Engine) {

	// v1 group route
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello Gohub!",
			})
		})
	}
}
