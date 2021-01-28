package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("App started in environment: ", os.Getenv("TIER"))
}
