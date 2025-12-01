package repositories

import "errors"

// ErrRecordNotFound is returned when no records are found
var ErrRecordNotFound = errors.New("record not found")
var ErrUserNotFound = errors.New("user not found")
