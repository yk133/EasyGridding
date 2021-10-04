package http

import (
	"EasyGridding/endpoint/log"
	"EasyGridding/endpoint/models"
	"EasyGridding/endpoint/proto"
	"EasyGridding/endpoint/tool"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var RecordDBService * models.RecordDBService
func NewRecordDBService(rdb * models.RecordDBService)   {
	RecordDBService = rdb
}

func HandlerGetHistory(g *gin.Context) {
	ctx := context.Background()
	req := proto.HandlerGetHistoryReq{}
	if err := g.ShouldBind(&req); err != nil {
		log.Log.Printf("HandlerGetHistoryReq binding error: %v", err)
		return
	}
	rc := RecordDBService

	var filter []*tool.Filter

	filter = append(filter, &tool.Filter{
		Name:   "userId",
		Values: []string{fmt.Sprintf("%d", req.UserId)}})

	filter = append(filter, &tool.Filter{
		Name:   "startTime",
		Values: []string{fmt.Sprintf("%d", req.StartTime)}}, &tool.Filter{
		Name:   "endTime",
		Values: []string{fmt.Sprintf("%d", req.EndTime)}})

	data, err := rc.GetRecordList(ctx, filter, req.Offset, req.Limit)
	if err != nil {
		log.Log.Printf("rc.GetRecordList failed %+v", err)
		FailResponse(g, http.StatusOK, err)
		return
	}

	SuccResponse(g, data)
	return
}

func HandlerCreateRecord(g *gin.Context) {
	ctx := context.Background()
	req := proto.HandlerCreateRecordReq{}
	if err := g.ShouldBind(&req); err != nil {
		log.Log.Printf("HandlerCreateRecordReq binding error: %v", err)
		return
	}
	rc := RecordDBService
	r := models.Record{
		UserId:    req.UserId,
		Type:      req.Type,
		Name:      req.Name,
		BuyPrice:  req.BuyPrice,
		WantPrice: req.WantPrice,
		Status:    req.Status,
	}

	err := rc.CreateRecord(ctx, &r)
	if err != nil {
		log.Log.Printf("rc.CreateRecord failed %+v", err)
		FailResponse(g, http.StatusOK, err)
		return
	}

	SuccResponse(g, "ok")
	return
}
