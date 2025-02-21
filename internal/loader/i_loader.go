package loader

import "database/sql"

type Loader interface {
	Load(db *sql.DB) error
}
