package postgresql

import (
	std_errors "errors"
)

/*
Each storage engine defines its own errors separately.
*/
var (
	ErrImplementMe        = std_errors.New("implement me")
	ErrRecordDoesNotExist = std_errors.New("record does not exist")
)
