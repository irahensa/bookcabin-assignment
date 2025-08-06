package usecase

import (
	"github.com/irahensa/bookcabin-assignment/backend/resource"
)

type IResource interface {
	GetAircraftByType(id int64) (res resource.Aircraft, err error)
}
