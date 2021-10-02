package log

import (
	"fmt"
	"log"
	"os"
)

var Log *log.Logger

func InitLog(file string) error {
	logfile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		return err
	}
	//defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("log started")
	Log = logger

	return nil
}
