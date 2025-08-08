package resource

import (
	"context"
	"database/sql"
	"time"

	"github.com/irahensa/bookcabin-assignment/backend/config"
	// logger "github.com/irahensa/bookcabin-assignment/backend/lib/log"
	_ "github.com/mattn/go-sqlite3"
)

func InitResource() (Resource, error) {
	db, err := connectDB()
	if err != nil {
		return Resource{}, err
	}

	return Resource{
		DB: db,
	}, nil
}

func connectDB() (*sql.DB, error) {
	cfg := config.GetConfig()

	db, err := sql.Open(cfg.Database.Driver, cfg.Database.ConnString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r Resource) GetAircraftByType(ctx context.Context, acType string) (res Aircraft, err error) {
	sqlStatement := `
	SELECT id, type, num_of_rows, row_arrangement FROM Aircrafts
	WHERE type = ?;`

	row := r.DB.QueryRow(sqlStatement, acType)
	err = row.Scan(&res.ID, &res.Type, &res.NumOfRows, &res.RowArrangement)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Resource) GetVoucherByFlightNumberAndDate(ctx context.Context, flightNumber string, date string) (res Voucher, err error) {
	sqlStatement := `
	SELECT id, crew_name, crew_id, flight_number, flight_date, aircraft_type, seats, created_at FROM Vouchers
	WHERE flight_number = ? AND flight_date = ?;`

	row := r.DB.QueryRow(sqlStatement, flightNumber, date)
	err = row.Scan(&res.ID, &res.CrewName, &res.CrewID, &res.FlightNumber, &res.FlightDate, &res.AircraftType, &res.Seats, &res.CreatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Resource) CreateVoucher(ctx context.Context, voucher Voucher) (err error) {
	sqlStatement := `
	INSERT INTO Vouchers (crew_name,crew_id,flight_number,flight_date,aircraft_type,seats,created_at)
	VALUES (?,?,?,?,?,?,?);`

	_, err = r.DB.Exec(sqlStatement, voucher.CrewName, voucher.CrewID, voucher.FlightNumber, voucher.FlightDate, voucher.AircraftType, voucher.Seats, time.Now())
	if err != nil {
		// logger.Error(err)
		return err
	}

	return nil
}
