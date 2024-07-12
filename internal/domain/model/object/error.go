package domain

import (
	"errors"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

var ERRORS []error = []error{
	ErrorNone,
	ErrorUserNotFound,
	ErrorUnauthorized,
	ErrorPageNotFound,
}

const (
	ErrUser         string = "user"
	ErrUnauthorized string = "unauthorized"
	ErrPage         string = "page"
)

var (
	ErrorNone         error = nil
	ErrorUnauthorized error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUnauthorized)
	ErrorUserNotFound error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + smodel.ErrNotFound)
	ErrorPageNotFound error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrPage + smodel.ErrSep + smodel.ErrNotFound)
)

func GetErrors() []error {
	return ERRORS
}
