package handler

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddAdmin_UserExists(t *testing.T) {
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
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	// Call the AddAdmin function
	err = handler.AddAdmin("admin", "admin", "admin")

	// Assert that the correct error is returned
	assert.Error(t, err)
	assert.EqualError(t, err, "data with email 'admin' already exists")

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddAdmin_UserDoesNotExist(t *testing.T) {
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
		WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	// Mock the INSERT query to succeed
	mock.ExpectExec("INSERT INTO users").
		WithArgs("new_admin", sqlmock.AnyArg(), "test@example.com", "admin", 0.0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the AddAdmin function
	err = handler.AddAdmin("new_admin", "password", "test@example.com")

	// Assert that no error occurred
	assert.NoError(t, err)

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
