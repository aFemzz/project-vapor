package handler

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"vapor/entity"
)

func TestCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	user := entity.User{Username: "test_user"}

	rows := sqlmock.NewRows([]string{"order_id", "title", "price"}).
		AddRow(1, "Game 1", 10.0).
		AddRow(1, "Game 2", 15.0)

	mock.ExpectQuery("SELECT od.order_id, g.title, g.price FROM users u JOIN orders o (.+) JOIN order_details od (.+) JOIN games g (.+)").
		WithArgs(user.Username, 0).
		WillReturnRows(rows)

	games, totalPrice, orderId, err := s.Cart(user)

	assert.NoError(t, err)
	assert.Len(t, games, 2)
	assert.Equal(t, 25.0, totalPrice)
	assert.Equal(t, 1, orderId)
}

func TestCart_ErrorQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	user := entity.User{Username: "test_user"}

	mock.ExpectQuery("SELECT od.order_id, g.title, g.price FROM users u JOIN orders o (.+) JOIN order_details od (.+) JOIN games g (.+)").
		WithArgs(user.Username, 0).
		WillReturnError(errors.New("database error"))

	_, _, _, err = s.Cart(user)

	assert.EqualError(t, err, "error while get data")
}
