// main_test.go

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var usersResponse []User
	err = json.NewDecoder(resp.Body).Decode(&usersResponse)
	assert.NoError(t, err)

	// Add your assertions for the response data here
}

func TestAddUser(t *testing.T) {
	app := setupApp()

	newUser := User{
		ID:       4,
		Username: "user4",
		Email:    "user4@example.com",
	}

	newUserJSON, err := json.Marshal(newUser)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/users/add", bytes.NewBuffer(newUserJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var addedUser User
	err = json.NewDecoder(resp.Body).Decode(&addedUser)
	assert.NoError(t, err)

	// Add your assertions for the added user here
}

func setupApp() *fiber.App {
	app := fiber.New()

	// Add routes and middleware here
	app.Get("/api/users", getUsers)
	app.Post("/api/users/add", addUser)

	return app
}

func TestMain(m *testing.M) {
	// Run tests and start your application for testing purposes if needed
	// For example, you can set up a test database or other dependencies here.
	code := m.Run()

	// Cleanup after testing if needed

	// Exit with the status code from tests
	os.Exit(code)
}
