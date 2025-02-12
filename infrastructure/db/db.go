package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/luisantonisu/wave15-grupo4/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

// user:password@tcp(localhost:3306)/dbname
func ConnectDB(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}
