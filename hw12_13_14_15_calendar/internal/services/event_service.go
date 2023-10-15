package services

import (
	"context"
	"time"

	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/dto"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type EventServiceInterface interface {
	Add(ctx context.Context, dto dto.CreateEventDto) (*storage.Event, error)
	Update(ctx context.Context, dto dto.UpdateEventDto) (*storage.Event, error)
	Delete(ctx context.Context, dto dto.DeleteEventDto) error
	ListOnDate(ctx context.Context, dto dto.ListByIntervalDto) ([]storage.Event, error)
	ListOnWeek(ctx context.Context, dto dto.ListByIntervalDto) ([]storage.Event, error)
	ListOnMonth(ctx context.Context, dto dto.ListByIntervalDto) ([]storage.Event, error)
}

type EventService struct {
	storage storage.Interface
}

func NewEventService(storage storage.Interface) EventServiceInterface {
	return &EventService{
		storage: storage,
	}
}

func (s EventService) Add(ctx context.Context, dto dto.CreateEventDto) (*storage.Event, error) {
	e := dto.ToModel()
	return s.storage.Add(ctx, e)
}

func (s EventService) Update(ctx context.Context, dto dto.UpdateEventDto) (*storage.Event, error) {
	um := dto.ToModel()
	return s.storage.Update(ctx, um)
}

func (s EventService) Delete(ctx context.Context, dto dto.DeleteEventDto) error {
	return s.storage.Delete(ctx, dto.UUID)
}

func (s EventService) ListOnDate(ctx context.Context, dto dto.ListByIntervalDto) ([]storage.Event, error) {
	return s.storage.ListByInterval(ctx, dto.Date, dto.Date.Add(1))
}

func (s EventService) ListOnWeek(ctx context.Context, dto dto.ListByIntervalDto) ([]storage.Event, error) {
	return s.storage.ListByInterval(ctx, dto.Date, dto.Date.Add(time.Hour*24*7))
}

func (s EventService) ListOnMonth(ctx context.Context, dto dto.ListByIntervalDto) ([]storage.Event, error) {
	return s.storage.ListByInterval(ctx, dto.Date, dto.Date.Add(time.Hour*24*30))
}
