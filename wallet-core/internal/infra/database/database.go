package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	User     string
	Password string
	Schema   string
	Port     string
	Options  []string
}

func NewDatabase(ctx context.Context) (*sql.DB, error) {
	dbUser := "root"
	dbPassword := "root"
	dbHost := "database"
	dbPort := "3306"
	dbSchema := "wallet"
	dbOptions := strings.Join([]string{"parseTime=true"}, "&")

	log.Println("connect database...")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbUser, dbPassword, dbHost, dbPort, dbSchema, dbOptions))
	if err != nil {
		return nil, err
	}

	log.Println("ping database...")
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
