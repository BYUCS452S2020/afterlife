package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/byuoitav/afterlife/handlers"
	"github.com/byuoitav/afterlife/pg"
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

		dbUsername string
		dbName     string
	)

	pflag.IntVarP(&port, "port", "P", 8080, "port to run the server on")
	pflag.Int8VarP(&logLevel, "log-level", "L", 0, "level to log at. refer to https://godoc.org/go.uber.org/zap/zapcore#Level for options")

	pflag.StringVar(&dbUsername, "db-username", "", "username to connect to postgres with")
	pflag.StringVar(&dbName, "db-name", "afterlife", "name of the postgres database to use")
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

	log, err := config.Build()
	if err != nil {
		fmt.Printf("unable to build logger: %s", err)
		os.Exit(1)
	}

	pgData, err := pg.New(dbUsername, dbName)
	if err != nil {
		log.Fatal("unable to connect to database", zap.Error(err))
	}

	handlers := handlers.Handlers{
		Logger:      log,
		DataService: pgData,
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

	api := e.Group("/api")

	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)
	api.POST("/logout", handlers.Logout)

	api.GET("/user", handlers.User)
	api.GET("/timeline", handlers.Timeline)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("unable to bind listener", zap.Error(err))
	}

	log.Info("Starting server", zap.String("on", lis.Addr().String()))
	err = e.Server.Serve(lis)
	switch {
	case errors.Is(err, http.ErrServerClosed):
	case err != nil:
		log.Fatal("failed to serve", zap.Error(err))
	}
}
