package payin

import (
	std_errors "errors"
)

/*
Each domain handler translates usecase errors to client level errors and
exposes them by following definition.
*/
var (
	ErrInternalServerError = std_errors.New("internal server error")
	ErrNotFound            = std_errors.New("not found")
)
