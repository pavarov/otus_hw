package eventhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/dto"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http/responses"
)

func (h *Handler) ListOnDate(c echo.Context) error {
	r, err := h.listByDateRequestValue(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	d := dto.ListByIntervalDto{Date: *r}

	l, err := h.eventService.ListOnDate(c.Request().Context(), d)
	if err != nil {
		h.log.Error("failed to get list of events", "error", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	resp := make([]responses.EventDataResponse, len(l))
	for i, event := range l {
		resp[i] = responses.EventDataResponse(event)
	}
	return c.JSON(http.StatusOK, resp)
}
