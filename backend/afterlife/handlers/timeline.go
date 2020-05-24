package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (h *Handlers) Timeline(c echo.Context) error {
	cookie, err := c.Cookie("afterlife-token")
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	timeline, err := h.DataService.Timeline(ctx, cookie.Value)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, timeline)
}
