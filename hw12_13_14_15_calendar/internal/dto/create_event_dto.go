package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type CreateEventDto struct {
	Title            string        `json:"title"`
	Start            time.Time     `json:"start"`
	End              time.Time     `json:"end"`
	Description      string        `json:"description"`
	UserID           uuid.UUID     `json:"userId"`
	NotificationTime time.Duration `json:"notificationTime"`
}

func (d *CreateEventDto) ToModel() storage.Event {
	return storage.Event{
		Title:            d.Title,
		Start:            d.Start,
		End:              d.End,
		Description:      d.Description,
		UserID:           d.UserID,
		NotificationTime: d.NotificationTime,
	}
}
