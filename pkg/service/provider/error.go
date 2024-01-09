package provider

import (
	std_errors "errors"
)

/*
Each domain defines its own errors separately.
*/
var (
	ErrProviderNotFound = std_errors.New("provider not found")
)
