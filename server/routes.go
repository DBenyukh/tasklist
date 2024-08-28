package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"tasklist/repo"
)

func Routes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Главная")
	})

	e.GET("/tasks", repo.GetAllTasks)
	e.POST("/tasks", repo.CreateTask)
	e.GET("/tasks/:id", repo.GetTask)
	e.PUT("/tasks/:id", repo.UpdateTask)
	e.DELETE("/tasks/:id", repo.DeleteTask)

}
