package handler

import "github.com/irahensa/bookcabin-assignment/backend/resource"

type Handler struct {
	Usecase IUsecase
}

type IUsecase interface {
	CheckVoucher(flightNumber, date string) (isExist bool, err error)
	GenerateVoucher(param resource.Voucher) (resp []string, err error)
}

type CheckVoucherParam struct {
	FlightNumber string `json:"flightNumber"`
	Date         string `json:"date"`
}

type CheckVoucherResponse struct {
	Exist bool `json:"exists"`
}

type GenerateVoucherParam struct {
	CrewName     string `json:"name"`
	CrewID       string `json:"id"`
	FlightNumber string `json:"flightNumber"`
	FlightDate   string `json:"date"`
	AircraftType string `json:"aircraft"`
}

type GenerateVoucherResponse struct {
	IsSuccess bool     `json:"success"`
	Seats     []string `json:"seats"`
}
