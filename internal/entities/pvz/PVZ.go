package pvz

import (
	"database/sql"
	"time"
)

type PVZ struct {
	ID               string         `db:"id" json:"id"`
	RegistrationDate time.Time      `db:"registration_date" json:"registration_date"`
	City             string         `db:"city" json:"city"`
	ActiveReception  sql.NullString `db:"active_reception" json:"active_reception"`
}
