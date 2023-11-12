package storage

import "errors"

var (
	ErrDateBusy             = errors.New("date is busy")
	ErrEventNotFound        = errors.New("event not found")
	ErrUpdateNoAffectedRows = errors.New("no affected rows while to update event")
)
