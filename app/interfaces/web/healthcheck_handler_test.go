package web

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const jsonHealthCheck = "{\"status\":\"UP\",\"version\":\"\"}\n"

func TestNewHealthHandler(t *testing.T) {

	t.Run("When response Healthcheck OK", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		NewHealthHandler(e)
		healthcheck := healthHandler{}

		healthcheck.HealthCheck(echoContext)

		assert.Equal(t, jsonHealthCheck, rec.Body.String())
		assert.Equal(t, http.StatusOK, rec.Code)
	})

}
