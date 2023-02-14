package handlers

import (
	"context"
	"time"

	"github.com/Rohan3011/go-todo-app/internal/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        int       `json:"_id"   bson:"_id"`
	Title     string    `json:"title" bson:"title"`
	Body      string    `json:"body"  bson:"body"`
	Done      bool      `json:"done"  bson:"done" `
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

func createTodo(c *fiber.Ctx) error {
	todo := Todo{
		ID:        primitive.NewObjectID().Timestamp().Minute(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.Database).Collection(string(db.TodosCollection))

	_, err = collection.InsertOne(context.TODO(), todo)

	if err != nil {
		return err
	}
	return c.JSON(todo)
}
