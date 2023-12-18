package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {

	// new a gin engine
	router := gin.New()

	// init route binding and middleware...
	bootstrap.SetupRoute(router)

	// run
	err := router.Run(":3000")
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}
}
