package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/hello-world/apps/user"
	"github.com/luyasr/hello-world/apps/user/api"
)

func UserRouterInit(router *gin.RouterGroup) {
	var svc user.Service
	handler := api.NewHandler(svc)
	u := router.Group("user")
	{
		u.GET(":id", handler.HelloUser)
		u.POST("", handler.CreateUser)
	}
}
