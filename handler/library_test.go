package handler

import (
	"errors"
	"testing"
	"vapor/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLibrary(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	user := entity.User{User_ID: 1}

	rows := sqlmock.NewRows([]string{"title"}).
		AddRow("Game 1").
		AddRow("Game 2")

	mock.ExpectQuery("SELECT g.title FROM games g JOIN order_details od (.+) JOIN orders o (.+)").
		WithArgs(user.User_ID, true).
		WillReturnRows(rows)

	data, err := s.Library(user)

	assert.NoError(t, err)
	expectedData := []string{"Game 1", "Game 2"}
	assert.Equal(t, expectedData, data)
}

func TestLibrary_ErrorQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	user := entity.User{User_ID: 1}

	mock.ExpectQuery("SELECT g.title FROM games g JOIN order_details od (.+) JOIN orders o (.+)").
		WithArgs(user.User_ID, true).
		WillReturnError(errors.New("database error"))

	_, err = s.Library(user)

	assert.EqualError(t, err, "error while get data")
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}
