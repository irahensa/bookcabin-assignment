package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/irahensa/bookcabin-assignment/backend/resource"
)

func InitUsecase(res IResource) UseCase {
	return UseCase{
		Resource: res,
	}
}

func (uc UseCase) CheckVoucher(flightNumber, date string) (isExist bool, err error) {
	resp, err := uc.Resource.GetVoucherByFlightNumberAndDate(flightNumber, date)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return false, nil
		}

		return false, err
	}

	if resp.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (uc UseCase) GenerateVoucher(param resource.Voucher) (resp []string, err error) {
	isExist, err := uc.CheckVoucher(param.FlightNumber, param.FlightDate)
	if err != nil {
		return resp, err
	}

	if isExist {
		return []string{}, errors.New("voucher already exist for this flight")
	}

	ac, err := uc.Resource.GetAircraftByType(param.AircraftType)
	if err != nil {
		return resp, err
	}

	m := make(map[string]bool)
	ss := strings.Split(ac.RowArrangement, ",")

	for i := 0; i < 3; i++ {
		row := rand.IntN(ac.NumOfRows) + 1
		col := rand.IntN(len(ss))
		seat := fmt.Sprintf("%d%s", row, ss[col])

		if !m[seat] {
			m[seat] = true
			resp = append(resp, seat)
		}
	}

	param.Seats = strings.Join(resp, ",")
	err = uc.Resource.CreateVoucher(param)
	if err != nil {
		return []string{}, err
	}

	return resp, nil
}
