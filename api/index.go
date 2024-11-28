package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	server "github.com/pichead/go-llm-chat/fiber"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := server.FiberApp()

	return adaptor.FiberApp(app)
}
