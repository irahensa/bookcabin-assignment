package usecase

import (
	"github.com/irahensa/bookcabin-assignment/backend/resource"
)

type IResource interface {
	GetAircraftByType(aircraftType string) (res resource.Aircraft, err error)
	GetVoucherByFlightNumberAndDate(flightNumber string, date string) (res resource.Voucher, err error)
	CreateVoucher(voucher resource.Voucher) (err error)
}

type UseCase struct {
	Resource IResource
}
