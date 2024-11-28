package main

import (
	"log"
	server "github.com/pichead/go-llm-chat/fiber"
)

func main() {

	app := server.FiberApp()
	log.Fatal(app.Listen(":3333"))

}
