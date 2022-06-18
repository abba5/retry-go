package retry_test

import (
	"errors"
	"testing"

	"github.com/abba5/retry-go/retry"
	"github.com/stretchr/testify/assert"
)

func Test_DoTryWithErrorMethod(t *testing.T) {
	attempt := 5
	count := 0
	function := func() error {
		count++
		return errors.New("some error")
	}

	err := retry.Do(attempt, function)
	assert.Error(t, err)
	assert.Equal(t, attempt, count)
	assert.Equal(t, "error limit exceded", err.Error())
}

func Test_DoWithSuccessMethod(t *testing.T) {
	attempt := 5
	count := 0
	function := func() error {
		count++
		if count == 4 {
			return nil
		}
		return errors.New("some error")
	}

	err := retry.Do(attempt, function)
	assert.Nil(t, err)
	assert.Equal(t, 4, count)
}

func Test_DoCWithRetryAlwaysMethod(t *testing.T) {
	attempt := 5
	count := 0
	function := func() error {
		count++
		return errors.New("some error")
	}
	shouldRetry := func(err error) bool {
		return true
	}

	err := retry.DoC(attempt, function, shouldRetry)
	assert.Error(t, err)
	assert.Equal(t, attempt, count)
	assert.Equal(t, "error limit exceded", err.Error())
}

func Test_DoCWitShouldNotRetry(t *testing.T) {
	attempt := 5
	count := 0
	function := func() error {
		count++
		return errors.New("some error")
	}
	shouldRetry := func(err error) bool {
		if count == 3 {
			return false
		}
		return true
	}

	err := retry.DoC(attempt, function, shouldRetry)
	assert.Error(t, err)
	assert.Equal(t, 3, count)
	assert.Equal(t, "some error", err.Error())
}
