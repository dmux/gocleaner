package internal

import (
	"log"
	"os"
)

var Logger *log.Logger
var LogFile *os.File

func SetupLogger() {
	var err error
	LogFile, err = os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Erro ao criar log.txt: %v", err)
	}
	Logger = log.New(LogFile, "", log.LstdFlags)
}
