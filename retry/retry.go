package retry

import "errors"

type Executor func() error
type ShouldRetry func(error) bool

var (
	maxRetryError = errors.New("error limit exceded")
)

func defaultDecider(err error) bool { return err != nil }

func DoC(maxAttempt int, fn Executor, decide ShouldRetry) error {
	count := 0
	for {
		err := fn()
		if err == nil {
			return nil
		}
		if !decide(err) {
			return err
		}
		count++
		if count > maxAttempt {
			return maxRetryError
		}
	}
}

func Do(maxAttempt int, fn Executor) error {
	return DoC(maxAttempt, fn, defaultDecider)
}
