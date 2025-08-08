package usecase

import (
	"context"

	"github.com/irahensa/bookcabin-assignment/backend/resource"
)

type IResource interface {
	GetAircraftByType(ctx context.Context, aircraftType string) (res resource.Aircraft, err error)
	GetVoucherByFlightNumberAndDate(ctx context.Context, flightNumber string, date string) (res resource.Voucher, err error)
	CreateVoucher(ctx context.Context, voucher resource.Voucher) (err error)
}

type UseCase struct {
	Resource IResource
}
