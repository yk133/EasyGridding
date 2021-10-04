package appcontext

import (
	"EasyGridding/endpoint/http"
	"EasyGridding/endpoint/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

type AppContext struct {
	MysqlConfig MysqlConfig `json:"mysql_config"`
	Run *Run `json:"run"`
	Token string `json:"token"`

	DB *gorm.DB
	Gin *gin.Engine
	RecordDBService *models.RecordDBService
}

type MysqlConfig struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Host string `json:"host"`
}

type Run struct{
	Port int `json:"port"`
}

var app *AppContext

func New() *AppContext {
	return &AppContext{}
}

func (a *AppContext) Init(file string) error {
	yamlFile, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		return err
	}

	return a.init()
}

func GetApp() *AppContext {
	return app
}

func (a *AppContext) init() error {
	a.Gin = gin.New()
	a.Gin.Use(middleware(a.Token))

	// DB
	db, err := gorm.Open(sqlite.Open("EasyGridding.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	a.DB = db

	// DBService
	a.RecordDBService = models.NewDBService(db,"record")

	// endpoint
	// Endpoint Init
	//
	http.NewRecordDBService(a.RecordDBService)

	http.AddFunc(a.Gin)
	return nil
}

func middleware(token string ) gin.HandlerFunc{

	return func(ctx *gin.Context) {
		tk:=ctx.GetHeader("token")
		if tk!=token {
			//ctx.Abort()
			log.Println("request token is invalid, token is ",tk)
			return
		}
	}
}