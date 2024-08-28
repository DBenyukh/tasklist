package repo

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"tasklist/env"
)

type Task struct {
}

const (
	tasksCollection = "tasks"
)

func GetAllTasks(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var tasks []Task
	defer cancel()

	conn, err := NewConn(env.GetDbUri())
	if err != nil {
		panic(err)
	}
	NewRepo(conn)
	results, err := conn.Collection(tasksCollection).Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleTask Task
		if err = results.Decode(&singleTask); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		tasks = append(tasks, singleTask)
	}

	return c.JSON(http.StatusOK, tasks)
}

func CreateTask(c echo.Context) error {
	return nil
}

func GetTask(c echo.Context) error {
	return nil
}

func UpdateTask(c echo.Context) error {
	return nil
}

func DeleteTask(c echo.Context) error {
	return nil
}
