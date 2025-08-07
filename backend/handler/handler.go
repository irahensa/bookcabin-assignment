package handler

import (
	"encoding/json"
	"net/http"

	"github.com/irahensa/bookcabin-assignment/backend/resource"
)

func InitHandler(uc IUsecase) Handler {
	return Handler{
		Usecase: uc,
	}
}

func (h *Handler) CheckVoucherHandler(w http.ResponseWriter, r *http.Request) {
	var param = CheckVoucherParam{}

	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Usecase.CheckVoucher(param.FlightNumber, param.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(CheckVoucherResponse{
		Exist: res,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) GenerateVoucherHandler(w http.ResponseWriter, r *http.Request) {
	var param = GenerateVoucherParam{}

	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Usecase.GenerateVoucher(resource.Voucher{
		CrewName:     param.CrewName,
		CrewID:       param.CrewID,
		FlightNumber: param.FlightNumber,
		FlightDate:   param.FlightDate,
		AircraftType: param.AircraftType,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(GenerateVoucherResponse{
		IsSuccess: true,
		Seats:     res,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
