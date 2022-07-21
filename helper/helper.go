package helper

import (
	"log"
	"os"
)

func WriteToLogFile(err error, errorCode int, message string) {
	file, logErr := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if logErr != nil {
		log.Fatal(logErr)
	}
	defer file.Close()
	log.SetOutput(file)
	if err != nil {
		log.Printf("%s, error model: %v, error code: %d", message, err, errorCode)
	} else {
		log.Printf("%s, error code: %d", message, errorCode)
	}
}
