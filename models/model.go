package models

import "errors"

var (
	ErrPersonNotNull  = errors.New("la persona no debe ser nula")
	ErrPersonNotExist = errors.New("la persona no existe")
)
