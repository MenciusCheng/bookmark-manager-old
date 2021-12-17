package http

import (
	"fmt"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/conf"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/service"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/util/logging"
	"github.com/gin-gonic/gin"
)

var (
	svc *service.Service

	httpServer *gin.Engine
)

// Init create a rpc server and run it
func Init(s *service.Service, conf *conf.Config) {
	svc = s

	// new http server
	httpServer = gin.New()

	// add namespace plugin
	httpServer.Use(gin.Logger())
	httpServer.Use(gin.Recovery())

	// register handler with http route
	initRoute(httpServer)

	// start a http server
	port := fmt.Sprintf(":%d", conf.Server.Port)
	go func() {
		if err := httpServer.Run(port); err != nil {
			logging.Error("http server start failed, err %v", err)
			panic(err)
		}
	}()

}

func Shutdown() {
	if httpServer != nil {
		//httpServer.Stop()
	}
	if svc != nil {
		svc.Close()
	}
}
