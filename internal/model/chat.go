package model

import (
	"database/sql"
	"time"
)

type Chat struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
