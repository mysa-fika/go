package user

import (
	"github.com/mysa-fika/go/domain/vocabulary"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateRequest struct {
	Realm    string
	Email    string
	Password string
}

func (cr *CreateRequest) Validate() error {
	return ozzo.ValidateStruct(
		cr,
		ozzo.Field(&cr.Email, ozzo.Required, is.Email),
		ozzo.Field(&cr.Password, ozzo.Required, ozzo.Length(8, 30)),
		ozzo.Field(&cr.Realm, ozzo.Required, ozzo.In(vocabulary.Client, vocabulary.Partner)),
	)
}
