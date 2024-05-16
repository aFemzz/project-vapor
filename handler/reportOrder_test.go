package handler

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestReportOrder_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	inputStart := "2024-01-01"
	inputEnd := "2024-12-31"

	rows := sqlmock.NewRows([]string{"title", "username", "date", "price"}).
		AddRow("Game 1", "user1", "2024-05-01", 10.0).
		AddRow("Game 2", "user2", "2024-05-02", 20.0)

	mock.ExpectQuery("SELECT g.title, u.username, od.date, g.price FROM users u JOIN orders o (.+) JOIN order_details od (.+) JOIN games g (.+)").
		WithArgs(true, inputStart, inputEnd).
		WillReturnRows(rows)

	expectedReports := []Report{
		{GameTitle: "Game 1", Username: "user1", Date: "2024-05-01"},
		{GameTitle: "Game 2", Username: "user2", Date: "2024-05-02"},
	}
	expectedTotalRevenue := 30.0

	reports, totalRevenue, err := s.ReportOrder(inputStart, inputEnd)

	assert.NoError(t, err)
	assert.Equal(t, expectedReports, reports)
	assert.Equal(t, expectedTotalRevenue, totalRevenue)
}

func TestReportOrder_ErrorQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	inputStart := "2024-01-01"
	inputEnd := "2024-12-31"

	mock.ExpectQuery("SELECT g.title, u.username, od.date, g.price FROM users u JOIN orders o (.+) JOIN order_details od (.+) JOIN games g (.+)").
		WithArgs(true, inputStart, inputEnd).
		WillReturnError(errors.New("database error"))

	_, _, err = s.ReportOrder(inputStart, inputEnd)

	assert.EqualError(t, err, "error while get data")
}
