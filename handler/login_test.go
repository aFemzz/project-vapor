package handler

import (
	"testing"
	"vapor/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Succeed(t *testing.T) {
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
	u := entity.User{}
	mock.ExpectQuery("^SELECT (.+) FROM users").
		WithArgs("yol@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "password", "user", "saldo"}).
			AddRow(4, "yol", "$2a$10$FwAiENlOg06fU/R.I3J0tuRnLusOjA9jGZD3TFAOAfVyQbEhG0cjq", "user", 0.0))

	// Call the Login function
	u, err = handler.Login("yol@gmail.com", "yol")

	// Assert that no error returned
	assert.NoError(t, err)
	assert.Equal(t, u.User_ID, 4)
	assert.Equal(t, u.Username, "yol")
	assert.Equal(t, u.Password, "$2a$10$FwAiENlOg06fU/R.I3J0tuRnLusOjA9jGZD3TFAOAfVyQbEhG0cjq")
	assert.Equal(t, u.Role, "user")
	assert.Equal(t, u.Saldo, 0.0)

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLogin_NoUser(t *testing.T) {
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
	u := entity.User{}
	mock.ExpectQuery("^SELECT (.+) FROM users").
		WithArgs("newyol@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "password", "user", "saldo"}))

	// Call the Login function
	u, err = handler.Login("newyol@gmail.com", "yol")

	// Assert that correct error return
	assert.Error(t, err)
	assert.EqualError(t, err, "password or user does not match")
	assert.Equal(t, u, entity.User{})

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
