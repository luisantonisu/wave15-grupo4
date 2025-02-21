package loader

import (
	"database/sql"
	"os"
)

func Load(db *sql.DB) error {
	file, err := os.ReadFile("infrastructure/db/data.sql")
	if err != nil {
		return err
	}
	sql := string(file)
	_, err = db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
