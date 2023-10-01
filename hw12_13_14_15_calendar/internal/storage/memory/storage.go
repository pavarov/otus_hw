package memorystorage

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Store struct {
	events sync.Map
}

func New() storage.StoreInterface {
	return &Store{}
}

func (s *Store) Add(_ context.Context, e storage.Event) error {
	s.events.Store(e.ID, e)
	return nil
}

func (s *Store) Find(_ context.Context, uuid uuid.UUID) (*storage.Event, error) {
	v, found := s.events.Load(uuid)
	if !found {
		return nil, storage.ErrEventNotFound
	}
	ev := v.(storage.Event)
	return &ev, nil
}

func (s *Store) Update(_ context.Context, e storage.Event) error {
	oldEvVal, loaded := s.events.Load(e.ID)
	if !loaded {
		return storage.ErrEventNotFound
	}
	oldEv := oldEvVal.(storage.Event)
	e.ID = oldEv.ID
	if oldEv != e {
		s.events.Store(e.ID, e)
	}

	return nil
}

func (s *Store) Delete(_ context.Context, uuid uuid.UUID) error {
	s.events.Delete(uuid)
	return nil
}

func (s *Store) List(_ context.Context) ([]storage.Event, error) {
	l := make([]storage.Event, 0)
	s.events.Range(func(key, value any) bool {
		l = append(l, value.(storage.Event))
		return true
	})
	return l, nil
}
