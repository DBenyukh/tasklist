package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"sync"
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

	}()
}
