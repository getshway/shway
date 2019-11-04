package domain

import "errors"

var (
	// ErrNotAvailableOS is error that the running OS is not allowed to run by shway
	ErrNotAvailableOS = errors.New("shway is able to run on MacOS only")
)
