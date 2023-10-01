package storage

import (
	"context"

	"github.com/google/uuid"
)

type StoreInterface interface {
	Add(ctx context.Context, e Event) error
	Update(ctx context.Context, e Event) error
	Delete(ctx context.Context, uuid uuid.UUID) error
	List(ctx context.Context) ([]Event, error)
	Find(ctx context.Context, uuid uuid.UUID) (*Event, error)
}
