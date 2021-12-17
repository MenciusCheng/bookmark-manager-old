package http

import (
	"context"
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

func getLinkList(c *gin.Context) {
	resp, err := svc.GetLinkList(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	okMsg := map[string]interface{}{"result": resp}
	c.JSON(http.StatusOK, okMsg)
}
