package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetDB() (*sql.DB, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME")))

    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    return db, err
} 