package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(router *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}
