package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/byuoitav/afterlife"
	"github.com/labstack/echo"
)

func (h *Handlers) Register(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	var req afterlife.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.DataService.Register(ctx, req); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handlers) Login(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	var req afterlife.LoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	tok, err := h.DataService.Login(ctx, req)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:    "afterlife-token",
		Value:   tok,
		Expires: time.Now().Add(24 * time.Hour),
	})

	return c.NoContent(http.StatusOK)
}

func (h *Handlers) Logout(c echo.Context) error {
	_, err := c.Cookie("afterlife-token")
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	//	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	//	defer cancel()

	// TODO delete the token

	c.SetCookie(&http.Cookie{
		Name:   "afterlife-token",
		Value:  "",
		MaxAge: -1,
	})

	return c.NoContent(http.StatusOK)
}

func (h *Handlers) User(c echo.Context) error {
	cookie, err := c.Cookie("afterlife-token")
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	user, err := h.DataService.User(ctx, cookie.Value)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
