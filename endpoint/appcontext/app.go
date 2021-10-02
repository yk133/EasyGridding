package appcontext

import (
	"EasyGridding/endpoint/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
)

type AppContext struct {
	MysqlConfig MysqlConfig `json:"mysql_config"`
	Run *Run `json:"run"`

	DB *gorm.DB
	Gin *gin.Engine
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

	db, err := gorm.Open(sqlite.Open("EasyGridding.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	a.DB = db

	http.AddFunc(a.Gin)

	return nil
}
