package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"sync"
	"tasklist/env"
)

func Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		e := echo.New()

		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		// Роуты
		Routes(e)

		// Старт сервера
		addr := env.GetWebAddr()
		if addr == "" {
			addr = ":8080"
		}

		if err := e.Start(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}
