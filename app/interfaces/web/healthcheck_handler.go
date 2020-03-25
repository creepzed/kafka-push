package web

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type healthHandler struct {
}

type health struct {
	Status string `json:"status"`
	Version string `json:"version"`
}

func NewHealthHandler(e *echo.Echo) {
	h := &healthHandler{
	}

	e.GET("/health", h.HealthCheck)
}

func (p *healthHandler) HealthCheck(c echo.Context) error {
	versionApp := os.Getenv("VERSION_APP")

	healthCheck := health{
		Status:  "UP",
		Version: versionApp,
	}

	return c.JSON(http.StatusOK, healthCheck)
}
