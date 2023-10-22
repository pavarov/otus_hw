package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Start time.Time `json:"start"`
	//nolint:tagliatelle
	UserID uuid.UUID `json:"user_id"`
}
