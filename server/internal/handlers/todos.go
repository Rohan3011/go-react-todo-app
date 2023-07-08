package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/Rohan3011/go-todo-app/internal/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Body      string             `json:"body"  bson:"body"`
	Done      bool               `json:"done"  bson:"done" `
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at"`
}

func GetTodo(c *fiber.Ctx) error {
	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.Database).Collection(string(db.TodosCollection))

	filter := bson.D{}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return err
	}

	var results []Todo
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return c.JSON(results)
}

func CreateTodo(c *fiber.Ctx) error {
	var todo Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

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

func UpdateTodo(c *fiber.Ctx) error {
	todoID := c.Params("id")

	// Parse the request body into a TodoUpdate struct
	var updateData struct {
		Done bool `json:"done"`
	}
	if err := c.BodyParser(&updateData); err != nil {
		return err
	}

	// Create an update filter
	filter := bson.M{"_id": todoID}

	// Create an update document with the new values
	update := bson.M{
		"$set": bson.M{
			"done":      updateData.Done,
			"updatedAt": time.Now(),
		},
	}

	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(db.Database).Collection(string(db.TodosCollection))

	// Perform the update operation
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	fmt.Println(result)

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo not found"})
	}

	return c.JSON(fiber.Map{"message": "Todo updated successfully"})
}

func DeleteTodo(c *fiber.Ctx) error {
	todoID := c.Params("id")

	filter := bson.M{"_id": todoID}

	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(db.Database).Collection(string(db.TodosCollection))

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo not found"})
	}

	return c.JSON(fiber.Map{"message": "Todo deleted successfully"})
}
