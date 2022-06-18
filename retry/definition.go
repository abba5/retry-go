package retry

type (
	Executor  func() error
	Retriable func(error) bool
)

func defaultRetriable(err error) bool {
	return err != nil
}
