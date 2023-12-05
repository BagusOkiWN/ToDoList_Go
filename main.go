package main


import (
    "net/http"

    "github.com/labstack/echo/v4"

)

type Task struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}

var tasks []Task

func main() {
    e := echo.New()

    e.GET("/tasks", getTasks)
    e.POST("/tasks", addTask)
    e.PUT("/tasks/:id", updateTask)
    e.DELETE("/tasks/:id", deleteTask)

    e.Start(":8080")
}

func getTasks(c echo.Context) error {
    // Implement logic to retrieve all tasks
    return c.JSON(http.StatusOK, tasks)
}

func addTask(c echo.Context) error {
    // Implement logic to add a new task
    return c.JSON(http.StatusCreated, tasks)
}

func updateTask(c echo.Context) error {
    // Implement logic to update the status of a task
    return c.NoContent(http.StatusNoContent)
}

func deleteTask(c echo.Context) error {
    // Implement logic to delete a task
    return c.NoContent(http.StatusNoContent)
}

func loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Implement logging logic (request method, path, duration)
        return next(c)
    }
}
