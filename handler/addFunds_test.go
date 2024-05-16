package handler

import (
	"testing"
	"vapor/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddFunds(t *testing.T) {
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

	// mock the update command
	mock.ExpectExec("UPDATE users").
		WithArgs(100.0, 10).
		WillReturnResult(sqlmock.NewResult(1, 1))

		// call the add funds
	err = handler.AddFunds(entity.User{User_ID: 10}, 100.0)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
