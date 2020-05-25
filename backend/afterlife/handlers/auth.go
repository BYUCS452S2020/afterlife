package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (h *Handlers) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("afterlife-token")
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
		defer cancel()

		user, err := h.DataService.User(ctx, cookie.Value)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		c.Set("user", user)
		return next(c)
	}
}
