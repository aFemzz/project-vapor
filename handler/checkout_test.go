package handler

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct {
	orderID    int
	totalPrice float64
	username   string
}

func TestCheckoutCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	orderID := 123
	totalPrice := 50.0
	username := "test_user"

	rows := sqlmock.NewRows([]string{"saldo"}).
		AddRow(100.0)

	mock.ExpectQuery("SELECT saldo FROM users").
		WithArgs(username).
		WillReturnRows(rows)

	mock.ExpectExec("UPDATE order_details SET is_purchased = 1").
		WithArgs(orderID, 0).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("UPDATE users SET saldo").
		WithArgs(totalPrice, username).
		WillReturnResult(sqlmock.NewResult(0, 1))

	newSaldo, err := s.CheckoutCart(orderID, totalPrice, username)

	assert.NoError(t, err)
	assert.Equal(t, 50.0, newSaldo)
}

func TestCheckoutCart_InsufficientBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	orderID := 123
	totalPrice := 150.0
	username := "test_user"

	rows := sqlmock.NewRows([]string{"saldo"}).
		AddRow(100.0)

	mock.ExpectQuery("SELECT saldo FROM users").
		WithArgs(username).
		WillReturnRows(rows)

	_, err = s.CheckoutCart(orderID, totalPrice, username)

	assert.EqualError(t, err, "you have no enough balance for this transaction, please add your funds")
}

func TestCheckoutCart_ErrorQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	orderID := 123
	totalPrice := 50.0
	username := "test_user"

	mock.ExpectQuery("SELECT saldo FROM users").
		WithArgs(username).
		WillReturnError(errors.New("database error"))

	_, err = s.CheckoutCart(orderID, totalPrice, username)

	assert.EqualError(t, err, "error while get data")
}
