package handler

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRegister_UserExist(t *testing.T) {
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
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs("yol@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

		// Call the AddAdmin function
	err = handler.Register("yol", "yol@gmail.com", "yol")

	// Assert that correct error returned
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("data with email '%s' already exists", "yol@gmail.com"))

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRegister_NewUser(t *testing.T) {
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
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs("user@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	// Mock the INSERT query to succeed
	mock.ExpectExec("INSERT INTO users").
		WithArgs("new_user", sqlmock.AnyArg(), "user@example.com", "user", 0.0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the Register function
	err = handler.Register("new_user", "user@example.com", "yol")

	// Assert that correct error returned
	assert.NoError(t, err)

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
