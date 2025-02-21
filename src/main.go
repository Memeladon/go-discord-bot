package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(2)
	}

	fmt.Println("Hello, world!")
}
