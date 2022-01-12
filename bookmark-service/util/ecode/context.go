package ecode

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// write response, error include business code and error msg
func JSON(c *gin.Context, data interface{}, err error) {
	w := NewWrapResp(data, err)
	c.JSON(w.Code, w)
}

func Success(c *gin.Context, data interface{}) {
	w := NewWrapResp(data, nil)
	c.JSON(http.StatusOK, w)
}

func Failure(c *gin.Context, data interface{}, err error) {
	w := NewWrapResp(data, err)
	c.JSON(w.Code, w)
}
