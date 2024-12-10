package main

import (
	"log"
 	"github.com/joho/godotenv"

	config "github.com/pichead/go-llm-chat/config"
	server "github.com/pichead/go-llm-chat/fiber"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		println("Load environment successfully")
	}

	app := server.FiberApp()

	println("port : ",config.Env.Port)
	log.Fatal(app.Listen(":" + config.Env.Port))

}
