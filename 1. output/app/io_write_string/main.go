package main

import (
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "Hello using io.WriteString!\n")
}
