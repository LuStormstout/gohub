// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {

	// register global middleware
	registerGlobalMiddleware(router)

	// register route
	routes.RegisterWebRoutes(router)

	// setting 404 handler
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(gin.Logger(), gin.Recovery())
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		// check accept header
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// return a string
			c.String(http.StatusNotFound, "404 Not Found")
		} else {
			// default return a json
			c.JSON(http.StatusNotFound, gin.H{
				"code": http.StatusNotFound,
				"msg":  "404 Not Found",
			})
		}
	})
}
