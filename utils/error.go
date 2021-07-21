package utils

import "errors"

var (
	ErrParameterInvalid = errors.New("Parameter valid")
	ErrExisted          = errors.New("record existed!")
	ErrNoExisted        = errors.New("Record no existed!")
	ErrFaild            = errors.New("Execute faild!")
)
