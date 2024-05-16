package handler_test

import (
	"testing"
	"vapor/entity"
	"vapor/handler"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestVaporWallet(t *testing.T) {
	// Create a mock SQL database.
	mockdb, mock, _ := sqlmock.New()

	// Create a handler with the mock SQL database.
	h := handler.Handler{DB: mockdb}

	// Prepare a user for testing.
	user := entity.User{User_ID: 1}

	// Mock the expected behavior of the database call.
	rows := sqlmock.NewRows([]string{"saldo"}).AddRow(100.0)
	mock.ExpectQuery("SELECT saldo FROM users WHERE user_id = ?").WithArgs(user.User_ID).WillReturnRows(rows)

	// Call the function being tested.
	wallet, err := h.VaporWallet(user)

	// Check if the error is as expected.
	assert.NoError(t, err, "Expected no error")

	// Check if the wallet balance is as expected.
	expectedWallet := 100.0 // Mocked wallet balance
	assert.Equal(t, expectedWallet, wallet, "Expected wallet balance to be 100.0")

	// Assert that all expectations were met.
	assert.NoError(t, mock.ExpectationsWereMet(), "Not all expectations were met")
}
