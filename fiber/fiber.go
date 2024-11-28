package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func FiberApp() *fiber.App {

	
	isProd := false

	app := fiber.New(fiber.Config{
		Prefork:       isProd,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Go LLM Chat App v0.0.4",
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(("hello world"))
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/books", getBooks)
	v1.Get("/books/:id", getBook)
	v1.Post("/books", createBook)
	v1.Put("/books/:id", updateBook)

	books = append(books, Book{ID: 1, Title: "HTML 101", Author: "Pichead"})
	books = append(books, Book{ID: 2, Title: "CSS 101", Author: "Maneerat"})
	return app

}

func getBooks(c *fiber.Ctx) error {

	return c.JSON(books)

}

func getBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books, *book)
	return c.JSON(c.Request().Body())
}

func updateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)

	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == id {
			books[i].Author = bookUpdate.Author
			books[i].Title = bookUpdate.Title
			return c.JSON(books[i])
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
