package requests

import "github.com/google/uuid"

type EventDelete struct {
	UUID uuid.UUID `json:"id" param:"id"`
}
