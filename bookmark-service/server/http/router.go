package http

import (
	"github.com/gin-gonic/gin"
)

func initRoute(r *gin.Engine) {
	r.GET("/ping", ping)

	r.POST("/link/create", createLink)
	//r.POST("/link/update", updateLink)
	//r.POST("/link/delete", deleteLink)
	r.GET("/link/list", getLinkList)
}
