package errorHandler

import "errors"

var ErrBookNotFound = "book not found"

var ErrBookIdMissing = "the book id does not exist or is checked out"

var ErrBookCheckedOut = "the book is checked out"

func NewError(message string) error {
	return errors.New(message)
}
