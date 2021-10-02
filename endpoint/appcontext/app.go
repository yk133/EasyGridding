package appcontext

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"os"
)

type AppContext struct {
	MysqlConfig MysqlConfig `json:"mysql_config"`

	Gin *gin.Engine

}
var app *AppContext
func (a *AppContext)ReadConf(file string ){
	f,err:=os.OpenFile(file,os.O_RDWR| os.O_APPEND,os.ModeAppend)
	if err!=nil{
		panic(err)
	}

	yaml.Marshal()
}

func GetApp()*AppContext{
	return app
}

type MysqlConfig struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Host string `json:"host"`
}

