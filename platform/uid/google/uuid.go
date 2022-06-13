package google

import (
	"login-app/platform/uid"

	"github.com/google/uuid"
)

type uuidgen struct{}

func New() uid.UUIDGen {
	return &uuidgen{}
}

func (u uuidgen) New() string {
	return uuid.NewString()
}
