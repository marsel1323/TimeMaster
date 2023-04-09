package handlers

import (
	"TimeMaster/internal/model"
	"TimeMaster/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type GoalHandler struct {
	goalService *service.GoalService
}

func NewGoalHandler(goalService *service.GoalService) *GoalHandler {
	return &GoalHandler{goalService: goalService}
}

func (h *GoalHandler) GetAllGoals(c echo.Context) error {
	goals, err := h.goalService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, goals)
}

func (h *GoalHandler) GetGoal(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	goal, err := h.goalService.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, goal)
}

func (h *GoalHandler) CreateGoal(c echo.Context) error {
	goal := new(model.Goal)
	if err := c.Bind(goal); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.goalService.Create(goal); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, goal)
}
