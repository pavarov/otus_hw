package eventhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/dto"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http/requests"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http/responses"
)

func (h *Handler) Create(c echo.Context) error {
	var r requests.EventCreateRequest
	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	d := dto.CreateEventDto(r)
	ev, err := h.eventService.Add(c.Request().Context(), d)
	if err != nil {
		h.log.Error("error while to add new event", "error", err)
		status := http.StatusInternalServerError
		return echo.NewHTTPError(status, http.StatusText(status))
	}

	resp := responses.EventDataResponse(*ev)
	return c.JSON(http.StatusCreated, resp)
}
