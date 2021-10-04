package http

import "github.com/gin-gonic/gin"


func AddFunc(r * gin.Engine){

	r.POST("/v1/easy_gridding/login", func(c *gin.Context) {

	})

	r.POST("/v1/easy_gridding/get_my_history", HandlerGetHistory)

	r.POST("/v1/easy_gridding/create_a_record",HandlerCreateRecord)

}
