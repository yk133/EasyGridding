package http

import (
	"EasyGridding/endpoint/log"
	"EasyGridding/endpoint/proto"
	"github.com/gin-gonic/gin"
)

func HandlerGetHistory(g *gin.Context){
	req:=proto.HandlerGetHistoryReq{}
	if err:=g.ShouldBind(&req);err!=nil{
		log.Log.Printf("HandlerGetHistoryReq binding error: %v", err)
		return
	}



	return
}
