package router

import (
	"douyinOrigin/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	apiServer := r.Group("/douyin")
	{
		//基础接口
		apiServer.POST("/user/register/", controller.Register)
		apiServer.POST("/user/login/", controller.Login)
		apiServer.GET("/user/", controller.UserInfo)
		apiServer.GET("/feed/", controller.Feed)
		apiServer.POST("/publish/action/", controller.Publish)
		apiServer.GET("/publish/list/", controller.PublishList)
	}
}
