package storage

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID               uuid.UUID     `db:"id"`
	Title            string        `db:"title"`
	Start            time.Time     `db:"start"`
	End              time.Time     `db:"end"`
	Description      string        `db:"description"`
	UserID           uuid.UUID     `db:"user_id"`
	NotificationTime time.Duration `db:"notification_time"`
}
