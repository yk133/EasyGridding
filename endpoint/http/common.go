package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccResponse(g* gin.Context,data interface{}){
	d:=struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Data: data,
	}
	g.JSON(http.StatusOK,&d)
}

func FailResponse(g* gin.Context,httpStatus int, err error) {
	g.JSON(httpStatus, gin.H{
		"code": "bad",
		"msg":  err.Error(),
	})
}
