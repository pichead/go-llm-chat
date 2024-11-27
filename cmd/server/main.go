package main

import (
	"log"

	"github.com/pichead/go-llm-chat/internal/app"
)

func main() {

	app := app.FiberApp()

	// รันเซิร์ฟเวอร์
	log.Fatal(app.Listen(":8000"))

}
