package payin

/*
Each domain defines its own errors separately.
*/

import (
	std_errors "errors"
)

var (
	ErrImplementMe     = std_errors.New("implement me")
	ErrPaymentNotFound = std_errors.New("payment not found")
)
