package main

import (
	"log"
	"os"

	"github.com/Rohan3011/go-todo-app/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

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

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/api/todos", handlers.GetTodo)
	app.Post("/api/todos", handlers.CreateTodo)
	app.Patch("/api/todos/:id", handlers.UpdateTodo)
	app.Delete("/api/todos/:id", handlers.DeleteTodo)

	app.Listen(":8080")
}
