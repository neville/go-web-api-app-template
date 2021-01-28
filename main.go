package main

import (
	"fmt"
	"os"

	"github.com/neville/go-web-api-app-template/server"
)

func main() {
	fmt.Println("App started in environment: ", os.Getenv("TIER"))

	server.Start()
}
