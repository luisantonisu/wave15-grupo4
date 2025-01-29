package error_handler

import "errors"

type MyError struct {
	Msg   string
	Error error
}

var CardNumberIdAlreadyInUse = errors.New("Card number id already in use")
var IDNotFound = errors.New("ID not found")
