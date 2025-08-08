package logger

import (
	"log"
)

func Error(err error) {
	log.Printf("ERROR: %v", err.Error())
}
