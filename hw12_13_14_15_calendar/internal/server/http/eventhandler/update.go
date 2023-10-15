package eventhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/dto"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http/requests"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http/responses"
)

func (h *Handler) Update(c echo.Context) error {
	var r requests.EventUpdateRequest
	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	d := dto.UpdateEventDto(r)
	ev, err := h.eventService.Update(c.Request().Context(), d)
	if err != nil {
		h.log.Error("error while to update event", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}

	resp := responses.EventDataResponse(*ev)
	return c.JSON(http.StatusNoContent, resp)
}
