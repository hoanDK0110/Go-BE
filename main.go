// main.go

package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{
	{ID: 1, Username: "user1", Email: "user1@example.com"},
	{ID: 2, Username: "user2", Email: "user2@example.com"},
	{ID: 3, Username: "user3", Email: "user3@example.com"},
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func addUser(c *fiber.Ctx) error {
	var newUser User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Generate a new unique ID for the new user
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	return c.JSON(newUser)
}

func main() {
	app := fiber.New()

	app.Get("/api/users", getUsers)
	app.Post("/api/users/add", addUser)

	app.Listen(":8080")
}
