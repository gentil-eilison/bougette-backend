package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthCheck struct {
	Healthy bool `json:"health"`
}

func (handler *Handler) HealthCheck(context echo.Context) error {
	healthCheckStruct := healthCheck{
		Healthy: true,
	}
	return context.JSON(http.StatusOK, &healthCheckStruct)
}
