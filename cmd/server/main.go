package main

import (
	"log"

	"github.com/pichead/go-llm-chat/pkg/app"
)

func main() {

	app := app.FiberApp()

	log.Fatal(app.Listen(":8000"))

}
