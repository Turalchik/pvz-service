package products

import (
	"database/sql"
	"time"
)

type Product struct {
	ID                string         `db:"id" json:"id"`
	ReceptionTime     time.Time      `db:"reception_time" json:"reception_time"`
	Type              string         `db:"type" json:"type"`
	ReceptionID       string         `db:"reception_id" json:"reception_id"`
	PreviousProductID sql.NullString `db:"previous_product_id" json:"previous_product_id"`
}
