package domain

import (
	"errors"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

var ERRORS []error = []error{
	ErrorNone,
	ErrorUserNotFound,
	ErrorUnauthorized,
}

const (
	ErrUser         string = "user"
	ErrUnauthorized string = "unauthorized"
)

var (
	ErrorNone         error = nil
	ErrorUnauthorized error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUnauthorized)
	ErrorUserNotFound error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + smodel.ErrNotFound)
)

func GetErrors() []error {
	return ERRORS
}
