package receptions

import (
	"database/sql"
	"time"
)

type Reception struct {
	ID            string         `db:"id" json:"id"`
	StartTime     time.Time      `db:"start_time" json:"start_time"`
	EndTime       sql.NullTime   `db:"end_time" json:"end_time"`
	PVZID         string         `db:"pvz_id" json:"pvz_id"`
	Status        string         `db:"status" json:"status"`
	LastProductID sql.NullString `db:"last_product_id" json:"last_product_id"`
}
