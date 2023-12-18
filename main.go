package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {

	// new a gin engine
	r := gin.New()

	// use middleware
	r.Use(gin.Logger(), gin.Recovery())

	// add a route
	r.GET("/", func(c *gin.Context) {
		// return a json
		c.JSON(http.StatusOK, gin.H{
			"test": "Hello Gohub!",
		})
	})

	// 404 handler
	r.NoRoute(func(c *gin.Context) {
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

	// run
	_ = r.Run(":8000")
}
