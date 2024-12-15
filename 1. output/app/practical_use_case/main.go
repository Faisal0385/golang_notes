package main

import (
	"fmt"
	"log"
)

func main() {
	errorCode := 404
	logMessage := fmt.Sprintf("Error occurred: Code %d", errorCode)
	log.Println(logMessage)
}
