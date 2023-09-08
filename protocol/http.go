package protocol

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/hello-world/protocol/router"
	"log"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	api := r.Group("apiV1")
	router.UserRouterInit(api)
	log.Fatal(r.Run())
}
