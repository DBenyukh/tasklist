package repo

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"tasklist/env"
	"tasklist/misc"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

const (
	tasksTable = "tasks"
)

func GetAllTasks(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := NewConn(env.GetDbUri())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer conn.Close()
	rep := NewRepo(conn)

	rows, err := rep.db.Query(ctx, "SELECT id, title, description, due_date, created_at, updated_at FROM "+tasksTable)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var dueDate, createdAt, updatedAt time.Time

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &dueDate, &createdAt, &updatedAt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// Преобразуем даты в формат RFC3339
		task.DueDate = misc.ToRFC3339(dueDate)
		task.CreatedAt = misc.ToRFC3339(createdAt)
		task.UpdatedAt = misc.ToRFC3339(updatedAt)

		tasks = append(tasks, task)
	}

	return c.JSON(http.StatusOK, tasks)
}

func CreateTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Парсим due_date из формата RFC3339
	dueDate, err := misc.ParseRFC3339(task.DueDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := NewConn(env.GetDbUri())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer conn.Close()
	rep := NewRepo(conn)

	now := time.Now()
	var id int
	err = rep.db.QueryRow(ctx, "INSERT INTO "+tasksTable+" (title, description, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		task.Title, task.Description, dueDate, now, now).Scan(&id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	task.ID = id
	task.CreatedAt = misc.ToRFC3339(now)
	task.UpdatedAt = misc.ToRFC3339(now)

	return c.JSON(http.StatusCreated, task)
}

func GetTask(c echo.Context) error {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := NewConn(env.GetDbUri())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer conn.Close()
	rep := NewRepo(conn)

	var task Task
	var dueDate, createdAt, updatedAt time.Time

	err = rep.db.QueryRow(ctx, "SELECT id, title, description, due_date, created_at, updated_at FROM "+tasksTable+" WHERE id = $1", id).
		Scan(&task.ID, &task.Title, &task.Description, &dueDate, &createdAt, &updatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.JSON(http.StatusNotFound, "Task not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	task.DueDate = misc.ToRFC3339(dueDate)
	task.CreatedAt = misc.ToRFC3339(createdAt)
	task.UpdatedAt = misc.ToRFC3339(updatedAt)

	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	id := c.Param("id")
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Парсим due_date из формата RFC3339
	dueDate, err := misc.ParseRFC3339(task.DueDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := NewConn(env.GetDbUri())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer conn.Close()
	rep := NewRepo(conn)

	now := time.Now()
	_, err = rep.db.Exec(ctx, "UPDATE "+tasksTable+" SET title = $1, description = $2, due_date = $3, updated_at = $4 WHERE id = $5",
		task.Title, task.Description, dueDate, now, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.JSON(http.StatusNotFound, "Task not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Извлекаем обновленную задачу из базы данных
	var createdAt, updatedAt time.Time

	row := rep.db.QueryRow(ctx, "SELECT id, title, description, due_date, created_at, updated_at FROM "+tasksTable+" WHERE id = $1", id)
	err = row.Scan(&task.ID, &task.Title, &task.Description, &dueDate, &createdAt, &updatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.JSON(http.StatusNotFound, "Task not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Преобразуем даты в формат RFC3339
	task.DueDate = misc.ToRFC3339(dueDate)
	task.CreatedAt = misc.ToRFC3339(createdAt)
	task.UpdatedAt = misc.ToRFC3339(updatedAt)

	return c.JSON(http.StatusOK, task)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := NewConn(env.GetDbUri())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer conn.Close()
	rep := NewRepo(conn)

	result, err := rep.db.Exec(ctx, "DELETE FROM "+tasksTable+" WHERE id = $1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Задача удалена"})
}
