package http

import (
	"context"
	"github.com/MenciusCheng/bookmark-manager/bookmark-service/util/ecode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(c *gin.Context) {
	if err := svc.Ping(context.Background()); err != nil {
		c.JSON(500, err)
		return
	}
	okMsg := map[string]string{"result": "ok"}
	c.JSON(200, okMsg)
}

func createLink(c *gin.Context) {
	if err := svc.Ping(context.Background()); err != nil {
		c.JSON(500, err)
		return
	}
	okMsg := map[string]string{"result": "ok"}
	c.JSON(200, okMsg)
}

// getLinkList godoc
// @summary 查询书签列表
// @description 查询书签列表
// @tags link
// @router /api/link/list [get]
// @success 200 {array} model.Link
func getLinkList(c *gin.Context) {
	resp, err := svc.GetLinkList(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	ecode.Success(c, resp)
}
