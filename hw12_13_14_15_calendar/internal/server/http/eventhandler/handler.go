package eventhandler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http/requests"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/services"
)

type Handler struct {
	eventService services.EventServiceInterface
	log          logger.Logger
}

func New(eventService services.EventServiceInterface, log logger.Logger) *Handler {
	return &Handler{
		eventService: eventService,
		log:          log,
	}
}

func (h *Handler) listByDateRequestValue(c echo.Context) (*time.Time, error) {
	var r requests.EventsByIntervalRequest
	err := c.Bind(&r)
	if err != nil {
		return nil, err
	}

	from, err := time.Parse(time.DateOnly, r.Date)
	if err != nil {
		return nil, err
	}

	return &from, nil
}
