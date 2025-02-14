package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/luisantonisu/wave15-grupo4/internal/config"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// user:password@tcp(localhost:3306)/dbname
func ConnectDB(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://infrastructure/db/migrations",
		"grupo4",
		driver,
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Println("Successfully connected to the database!")
	return db
}
