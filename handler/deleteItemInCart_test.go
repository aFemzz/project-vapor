package handler

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteItemInCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	orderID := 123
	gameTitle := "Game 1"

	rows := sqlmock.NewRows([]string{"order_detail_id"}).
		AddRow(1) // Simulate the item exists in the cart

	mock.ExpectQuery("SELECT od.order_detail_id FROM games (.+)").
		WithArgs(gameTitle, orderID).
		WillReturnRows(rows)

	mock.ExpectExec("DELETE FROM order_details").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = s.DeleteItemInCart(orderID, gameTitle)

	assert.NoError(t, err)
}
