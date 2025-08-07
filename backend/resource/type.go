package resource

import (
	"database/sql"
)

type Resource struct {
	DB *sql.DB
}

type Voucher struct {
	ID           int64  `json:"id"             db:"id"`
	CrewName     string `json:"crew_name"      db:"crew_name"`
	CrewID       string `json:"crew_id"        db:"crew_id"`
	FlightNumber string `json:"flight_number"  db:"flight_number"`
	FlightDate   string `json:"flight_date"    db:"flight_date"`
	AircraftType string `json:"aircraft_type"  db:"aircraft_type"`
	Seats        string `json:"seats"          db:"seats"`
	CreatedAt    string `json:"created_at"     db:"created_at"`
}

type Aircraft struct {
	ID             int64  `json:"id"               db:"id"`
	Type           string `json:"type"             db:"type"`
	NumOfRows      int    `json:"num_of_rows"      db:"num_of_rows"`
	RowArrangement string `json:"row_arrangement"  db:"row_arrangement"`
}
