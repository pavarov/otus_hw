package requests

import (
	"time"

	"github.com/google/uuid"
)

type EventCreateRequest struct {
	Title       string    `json:"title"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Description string    `json:"description"`
	//nolint:tagliatelle
	UserID uuid.UUID `json:"user_id"`
	//nolint:tagliatelle
	NotificationTime time.Duration `json:"notification_time"`
}
