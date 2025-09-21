package errorHandler

import "errors"

// TODO: Refactor code to have singular go module and generic folders in one location for the repo to reduce duplicating code
// TODO: Or breakout into a separate repo if the projects become too large

var ErrUserNotFound = "user not found"

func NewError(message string) error {
	return errors.New(message)
}
