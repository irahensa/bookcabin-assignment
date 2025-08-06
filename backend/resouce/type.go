package resource

import (
	"database/sql"
)

type Resource struct {
	DB *sql.DB
}
