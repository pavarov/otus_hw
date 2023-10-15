package responses

import (
	"time"

	"github.com/google/uuid"
)

type EventDataResponse struct {
	ID               uuid.UUID     `json:"id"`
	Title            string        `json:"title"`
	Start            time.Time     `json:"start"`
	End              time.Time     `json:"end"`
	Description      string        `json:"description"`
	UserID           uuid.UUID     `json:"userId"`
	NotificationTime time.Duration `json:"notificationTime"`
}
