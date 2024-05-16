package handler_test

import (
	"testing"
	"vapor/entity"
	"vapor/handler"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestPurchaseGame(t *testing.T) {
	// Create a mock SQL database.
	mockdb, mock, _ := sqlmock.New()

	// Create a handler with the mock SQL database.
	h := handler.Handler{DB: mockdb}

	// Prepare a user for testing.
	user := entity.User{User_ID: 13}

}
