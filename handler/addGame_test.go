package handler

import (
	"fmt"
	"testing"
	"vapor/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddGame_GameExists(t *testing.T) {
	// Create a new mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	// Create a new handler with the mock DB
	handler := &Handler{
		DB: mockDB,
	}

	// Mock the SELECT query to return false (user exists)
	var game entity.Games
	game.Title = "Epic Adventure"
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(game.Title).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	// Call the AddAdmin function

	err = handler.AddGame(game)

	// Assert that the correct error is returned
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("game with title '%s' already exists", game.Title))

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddGame_GameDoesNotExist(t *testing.T) {
	// Create a new mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	// Create a new handler with the mock DB
	handler := &Handler{
		DB: mockDB,
	}

	// Mock the SELECT query to return false (user does not exist)
	g := entity.Games{
		Title:       "Zuma",
		Description: "main kodok",
		Price:       10,
		Developer:   "-",
		Publisher:   "-",
		Rating:      1,
	}
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(g.Title).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	// Mock the INSERT query to succeed
	mock.ExpectExec("INSERT INTO games").
		WithArgs(g.Title, g.Description, g.Price, g.Developer, g.Publisher, g.Rating, false).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the AddAdmin function
	err = handler.AddGame(g)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
