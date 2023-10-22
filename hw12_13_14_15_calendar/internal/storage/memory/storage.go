package memorystorage

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	events sync.Map
}

func New() storage.Interface {
	return &Storage{}
}

func (s *Storage) Add(_ context.Context, e storage.Event) (*storage.Event, error) {
	s.events.Store(e.ID, e)

	ev, _ := s.events.Load(e.ID)
	re := ev.(storage.Event)
	return &re, nil
}

func (s *Storage) Find(_ context.Context, uuid uuid.UUID) (*storage.Event, error) {
	v, found := s.events.Load(uuid)
	if !found {
		return nil, storage.ErrEventNotFound
	}
	ev := v.(storage.Event)
	return &ev, nil
}

func (s *Storage) Update(_ context.Context, e storage.Event) (*storage.Event, error) {
	oldEvVal, loaded := s.events.Load(e.ID)
	if !loaded {
		return nil, storage.ErrEventNotFound
	}
	oldEv := oldEvVal.(storage.Event)
	e.ID = oldEv.ID
	if oldEv != e {
		s.events.Store(e.ID, e)
	}

	return &e, nil
}

func (s *Storage) Delete(_ context.Context, uuid uuid.UUID) error {
	s.events.Delete(uuid)
	return nil
}

func (s *Storage) ListByInterval(_ context.Context, _ time.Time, _ time.Time) ([]storage.Event, error) {
	l := make([]storage.Event, 0)
	s.events.Range(func(key, value any) bool {
		l = append(l, value.(storage.Event))
		return true
	})
	return l, nil
}

func (s *Storage) ListToNotify(_ context.Context) ([]storage.Event, error) {
	panic("implement me")
}

func (s *Storage) RemoveOld(_ context.Context, _ time.Time) error {
	panic("implement me")
}
