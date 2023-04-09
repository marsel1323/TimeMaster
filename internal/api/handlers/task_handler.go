package handlers

import (
	"TimeMaster/internal/model"
	"TimeMaster/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	goals, err := h.taskService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, goals)
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskService.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.taskService.Create(task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, task)
}
