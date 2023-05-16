package mlog

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
	Warning *log.Logger
)

func InitLogger() {
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLog := log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLog := log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = infoLog
	Error = errorLog
	Debug = debugLog
	Warning = warnLog
}
