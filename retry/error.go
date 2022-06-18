package retry

import "errors"

var (
	maxRetryError = errors.New("error limit exceded")
)
