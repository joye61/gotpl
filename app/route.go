package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joye61/gotpl/app/controller"
)

func RouterSet(r *gin.Engine) {
	// 每个控制器文件对应一个分组
	test := r.Group("/test", controller.TestIndex)
	{
		test.Use()
		test.GET("/v1", controller.TestShow)
	}
}
