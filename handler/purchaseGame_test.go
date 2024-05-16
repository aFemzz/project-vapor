package handler_test

import (
	"testing"
	"time"
	"vapor/entity"
	"vapor/handler"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPurchaseGame(t *testing.T) {
	// Create a mock SQL database.
	mockdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockdb.Close()

	// Create a handler with the mock SQL database.
	h := handler.Handler{DB: mockdb}

	// Prepare a user for testing.
	user := entity.User{User_ID: 1}
	gameID := entity.Games{GameID: 2}

	mock.ExpectExec("INSERT INTO order_details").
		WithArgs(user.User_ID, gameID.GameID, false, time.Now().Format("2006-01-02")).
		WillReturnResult(sqlmock.NewResult(1, 2))

	err = h.InsertToCart(1, 2)

	assert.NoError(t, err)

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
