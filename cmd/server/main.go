package main

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

func main() {

	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "HTML 101", Author: "Pichead"})
	books = append(books, Book{ID: 2, Title: "CSS 101", Author: "Maneerat"})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(("hello world"))
	})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)

	app.Listen(":8000")

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

	return c.JSON(books)
}
