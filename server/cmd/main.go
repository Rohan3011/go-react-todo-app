package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool   `json:"done"`
}

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CLIENT_URL"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)
		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}
		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}
		return c.JSON(todos)
	})

	app.Listen(":8080")
}