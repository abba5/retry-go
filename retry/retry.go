package retry

func DoC(maxAttempt int, fn Executor, retriable Retriable) error {
	count := 1
	for {
		err := fn()
		if err == nil {
			return nil
		}
		if !retriable(err) {
			return err
		}
		count++
		if count > maxAttempt {
			return maxRetryError
		}
	}
}

func Do(maxAttempt int, fn Executor) error {
	return DoC(maxAttempt, fn, defaultRetriable)
}
