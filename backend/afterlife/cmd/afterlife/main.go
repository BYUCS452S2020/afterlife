package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/byuoitav/afterlife/handlers"
	"github.com/byuoitav/afterlife/mock"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	var (
		port     int
		logLevel int8
	)

	pflag.IntVarP(&port, "port", "P", 8080, "port to run the server on")
	pflag.Int8VarP(&logLevel, "log-level", "L", 0, "level to log at. refer to https://godoc.org/go.uber.org/zap/zapcore#Level for options")
	pflag.Parse()

	// build the logger
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(logLevel)),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json", EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "@",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	plain, err := config.Build()
	if err != nil {
		fmt.Printf("unable to build logger: %s", err)
		os.Exit(1)
	}

	handlers := handlers.Handlers{
		Logger:      plain,
		DataService: &mock.DataService{},
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost},
		MaxAge:       7200, // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age
	}))

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})

	e.PUT("/login", handlers.Login)
	e.GET("/user", handlers.User)
	e.GET("/timeline", handlers.Timeline)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		plain.Fatal("unable to bind listener", zap.Error(err))
	}

	plain.Info("Starting server", zap.String("on", lis.Addr().String()))
	err = e.Server.Serve(lis)
	switch {
	case errors.Is(err, http.ErrServerClosed):
	case err != nil:
		plain.Fatal("failed to serve", zap.Error(err))
	}
}
