package main

import (
	"fmt"
	"log"
	"ot-go/httpserver"
)

func main() {
	port := ":8080"
	fmt.Println("Server starting on http://localhost" + port)
	if err := httpserver.InitServer(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
