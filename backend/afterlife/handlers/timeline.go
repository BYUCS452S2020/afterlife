package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/byuoitav/afterlife"
	"github.com/labstack/echo"
)

func (h *Handlers) CreateEvent(c echo.Context) error {
	user, ok := c.Get("user").(afterlife.User)
	if !ok {
		return c.String(http.StatusUnauthorized, "user missing in context")
	}

	var event afterlife.Event
	if err := c.Bind(&event); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	if err := h.DataService.CreateEvent(ctx, user.ID, event); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handlers) UpdateEvent(c echo.Context) error {
	user, ok := c.Get("user").(afterlife.User)
	if !ok {
		return c.String(http.StatusUnauthorized, "user missing in context")
	}

	var event afterlife.Event
	if err := c.Bind(&event); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	if err := h.DataService.UpdateEvent(ctx, user.ID, event); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handlers) DeleteEvent(c echo.Context) error {
	user, ok := c.Get("user").(afterlife.User)
	if !ok {
		return c.String(http.StatusUnauthorized, "user missing in context")
	}

	id := c.Param("id")
	if len(id) == 0 {
		return c.String(http.StatusBadRequest, "missing id")
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	if err := h.DataService.DeleteEvent(ctx, user.ID, afterlife.EventID(id)); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handlers) Timeline(c echo.Context) error {
	user, ok := c.Get("user").(afterlife.User)
	if !ok {
		return c.String(http.StatusUnauthorized, "user missing in context")
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	timeline, err := h.DataService.Timeline(ctx, user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, timeline)
}
