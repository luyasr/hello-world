package protocol

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/hello-world/config"
	"github.com/luyasr/hello-world/protocol/router"
	"log"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	api := r.Group("apiV1")
	router.UserRouterInit(api)
	log.Fatal(r.Run(fmt.Sprintf(":%d", config.Conf.Server.Port)))
}
