package main

import (
	"fmt"
	"log"
)

func main() {
	log.Fatal("A critical error occurred!")
	fmt.Println("This will not execute.")
}
