// Package errors - wrapping errors
package errors

import (
	"fmt"
	"strings"
)

// Wrap - wraps the previous error with current error/ message
func Wrap(previous error, current interface{}) error {
	var currentError error
	switch c := current.(type) {
	case error:
		currentError = c
	case string:
		if strings.TrimSpace(c) == "" {
			currentError = invalidWrap("empty string")
		} else {
			currentError = new(c)
		}
	default:
		currentError = invalidWrap("unsupported type, it should be either of type error/string")
	}

	return &withError{
		previous: previous,
		current:  currentError,
	}
}

func UnWrap(err error) (current error, previous error) {
	switch e := err.(type) {
	case *withError:
		return e.current, e.previous
	default:
		return e, nil
	}
}

func invalidWrap(err string) *Error {
	code := "incorrect_usage"
	msg := fmt.Sprintf("you should pass either error or message to properly wrap the error! - wrapped with %s", err)
	return new(code, msg)
}
