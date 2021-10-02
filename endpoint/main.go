package main

import (
	"EasyGridding/endpoint/appcontext"
	"EasyGridding/endpoint/log"
	"fmt"
)

func main(){

	app:=appcontext.New()

	log.InitLog("./log.log")
	err:=app.Init("./config.yaml")
	if err != nil {
		panic(err)
	}

	app.Gin.Run(fmt.Sprintf(":%d",app.Run.Port))

	return
}
