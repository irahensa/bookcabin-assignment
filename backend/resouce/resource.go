package resource

import (
	"database/sql"
	"fmt"

	"github.com/irahensa/bookcabin-assignment/backend/config"
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
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}

func (r Resource) GetAircraftByType(id int64) (res Aircraft, err error) {
	sqlStatement := `
	SELECT id, type, num_of_rows, row_arrangement FROM Users
	WHERE ID = ?;`

	row := r.DB.QueryRow(sqlStatement, id)
	err = row.Scan(&res.ID, &res.Type, &res.NumOfRows, &res.RowArrangement)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Resource) GetVoucherByFlightNumberAndDate(flightNumber string, date string) (res Voucher, err error) {
	sqlStatement := `
	SELECT id, crew_name, crew_id, flight_number, flight_date, aircraft_type, seats, created_at FROM Airplanes
	WHERE type = ?;`

	row := r.DB.QueryRow(sqlStatement, flightNumber, date)
	err = row.Scan(&res.ID, &res.CrewName, &res.CrewID, &res.FlightNumber, &res.FlightDate, &res.AircraftType, &res.Seats, &res.CreatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
