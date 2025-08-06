package resource

import (
	"database/sql"
	"fmt"

	"github.com/imam-rahensa/bookcabin-assignment/backend/config"
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

	connstring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Database.User, cfg.Database.Password, cfg.Database.Address, cfg.Database.Port, cfg.Database.DBName)
	db, err := sql.Open("mysql", connstring)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
