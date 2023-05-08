package mlog

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func InitLogger() {
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = infoLog
	Error = errorLog
}
