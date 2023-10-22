package storage

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Interface interface {
	Add(ctx context.Context, e Event) (*Event, error)
	Update(ctx context.Context, e Event) (*Event, error)
	Delete(ctx context.Context, uuid uuid.UUID) error
	ListByInterval(ctx context.Context, from time.Time, to time.Time) ([]Event, error)
	Find(ctx context.Context, uuid uuid.UUID) (*Event, error)
	ListToNotify(ctx context.Context) ([]Event, error)
	RemoveOld(ctx context.Context, from time.Time) error
}
