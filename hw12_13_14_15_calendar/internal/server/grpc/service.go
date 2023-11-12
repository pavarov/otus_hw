package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	dto2 "github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/dto"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/pb"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/services"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Service struct {
	pb.UnimplementedEventServiceServer
	service services.EventServiceInterface
}

func NewService(service services.EventServiceInterface) pb.EventServiceServer {
	return &Service{
		service: service,
	}
}

func (s *Service) Create(ctx context.Context, event *pb.Event) (*pb.EventDataResponse, error) {
	userID, err := uuid.Parse(event.GetUserId())
	if err != nil {
		return nil, err
	}
	dto := dto2.CreateEventDto{
		Title:            event.GetTitle(),
		Start:            event.GetStart().AsTime(),
		End:              event.GetEnd().AsTime(),
		Description:      event.GetDescription(),
		UserID:           userID,
		NotificationTime: event.GetNotifyTime().AsDuration(),
	}
	add, err := s.service.Add(ctx, dto)
	if err != nil {
		return nil, err
	}
	event.Id = add.ID.String()
	return &pb.EventDataResponse{Event: event}, nil
}

func (s *Service) Update(ctx context.Context, event *pb.Event) (*pb.EventDataResponse, error) {
	id, err := uuid.Parse(event.GetId())
	if err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(event.GetUserId())
	if err != nil {
		return nil, err
	}
	dto := dto2.UpdateEventDto{
		ID:               id,
		Title:            event.GetTitle(),
		Start:            event.GetStart().AsTime(),
		End:              event.GetEnd().AsTime(),
		Description:      event.GetDescription(),
		UserID:           userID,
		NotificationTime: event.GetNotifyTime().AsDuration(),
	}
	add, err := s.service.Update(ctx, dto)
	if err != nil {
		return nil, err
	}
	event.Id = add.ID.String()
	return &pb.EventDataResponse{Event: event}, nil
}

func (s *Service) Delete(ctx context.Context, request *pb.DeleteEventRequest) (*pb.EventDeleteResponse, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	dto := dto2.DeleteEventDto{
		UUID: id,
	}
	err = s.service.Delete(ctx, dto)
	if err != nil {
		return nil, err
	}
	return &pb.EventDeleteResponse{}, nil
}

func (s *Service) ListOnDate(ctx context.Context, interval *pb.ListByInterval) (*pb.ListByDateResponse, error) {
	dto := dto2.ListByIntervalDto{Date: interval.GetDate().AsTime()}
	list, err := s.service.ListOnDate(ctx, dto)
	if err != nil {
		return nil, err
	}
	return s.listEventsToResponse(list), nil
}

func (s *Service) ListOnWeek(ctx context.Context, interval *pb.ListByInterval) (*pb.ListByDateResponse, error) {
	dto := dto2.ListByIntervalDto{Date: interval.GetDate().AsTime()}
	list, err := s.service.ListOnDate(ctx, dto)
	if err != nil {
		return nil, err
	}
	return s.listEventsToResponse(list), nil
}

func (s *Service) ListOnMonth(ctx context.Context, interval *pb.ListByInterval) (*pb.ListByDateResponse, error) {
	dto := dto2.ListByIntervalDto{Date: interval.GetDate().AsTime()}
	list, err := s.service.ListOnDate(ctx, dto)
	if err != nil {
		return nil, err
	}
	return s.listEventsToResponse(list), nil
}

func (s *Service) listEventsToResponse(list []storage.Event) *pb.ListByDateResponse {
	response := make([]*pb.Event, 0, len(list))
	for _, event := range list {
		d := &event.Description
		response = append(response, &pb.Event{
			Id:          event.ID.String(),
			Title:       event.Title,
			Start:       &timestamp.Timestamp{Seconds: event.Start.Unix()},
			End:         &timestamp.Timestamp{Seconds: event.End.Unix()},
			Description: d,
			UserId:      event.UserID.String(),
			NotifyTime:  &duration.Duration{Seconds: int64(event.NotificationTime)},
		})
	}
	return &pb.ListByDateResponse{
		Events: response,
	}
}
