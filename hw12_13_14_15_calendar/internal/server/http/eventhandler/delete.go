package eventhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/dto"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/server/http/requests"
)

func (h *Handler) Delete(c echo.Context) error {
	var r requests.EventDelete
	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	d := dto.DeleteEventDto(r)

	err := h.eventService.Delete(c.Request().Context(), d)
	if err != nil {
		h.log.Error("failed to delete event", "error", err)
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
