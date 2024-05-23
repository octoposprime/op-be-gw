package domain

import (
	"errors"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

var ERRORS []error = []error{
	ErrorNone,
	ErrorUserNotFound,
	ErrorDlrNotFound,
}

const (
	ErrUser string = "user"
	ErrDlr  string = "dlr",
)

var (
	ErrorNone         error = nil
	ErrorUserNotFound error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + smodel.ErrNotFound)
	ErrorDlrNotFound  error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrDlr + smodel.ErrSep + smodel.ErrNotFound)
)

func GetErrors() []error {
	return ERRORS
}
