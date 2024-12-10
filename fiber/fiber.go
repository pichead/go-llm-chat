package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	bookController "github.com/pichead/go-llm-chat/controller/bookController"
)

func FiberApp() *fiber.App {
	isProd := false
	engine := html.New("./view", ".html")

	app := fiber.New(fiber.Config{
		Prefork:       isProd,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Go LLM Chat App v0.0.4",
		Views: engine,

	})

	// Setup route

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust this to be more restrictive if needed
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Static("/", "./public")
	app.Get("/html", renderTemplate)


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/upload", uploadImage)

	v1.Get("/books", bookController.GetBooks)
	v1.Get("/books/:id", bookController.GetBook)
	v1.Post("/books", bookController.CreateBook)
	v1.Put("/books/:id", bookController.UpdateBook)

	return app
}

func uploadImage(c *fiber.Ctx) error {
	// Read file from request
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Save the file to the server
	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File uploaded successfully: " + file.Filename)
}



func renderTemplate(c *fiber.Ctx) error {
	// Render the template with variable data
	return c.Render("template", fiber.Map{
	  "Name": "World",
	})
  }