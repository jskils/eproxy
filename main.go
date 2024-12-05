package main

import (
	"eproxy/router"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// !@#Aa40153907l
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	p := os.Getenv("PORT")
	r := router.SetupRouter()
	log.Println(fmt.Sprintf("Server is running on :%s", p))
	if err := r.Run(fmt.Sprintf(":%s", p)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
