package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/MenciusCheng/bookmark-manager/bookmark-service/docs"
)

func initRoute(r *gin.Engine) {
	r.GET("/api/ping", ping)

	r.POST("/api/link/create", createLink)
	//r.POST("/link/update", updateLink)
	//r.POST("/link/delete", deleteLink)
	r.GET("/api/link/list", getLinkList)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
