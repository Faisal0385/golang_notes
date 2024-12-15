package main

import (
	"fmt"
	"log"
)

func main() {
	log.Panic("This will panic!")
	fmt.Println("This will not execute.")
}
