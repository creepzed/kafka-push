package web

import (
	"errors"
	"fmt"
	"github.com/kafka-push/app/application/payload_usecase"
	"github.com/kafka-push/app/shared/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type payloadCreateController struct {
	useCase payload_usecase.PayloadUseCase
}

func NewPayloadController(e *echo.Echo, useCase payload_usecase.PayloadUseCase) *payloadCreateController {
	m := &payloadCreateController{
		useCase: useCase,
	}
	e.POST("/payload/:topic", m.PayloadPost)

	return m
}

func (p *payloadCreateController) PayloadPost(c echo.Context) error {
	topic := c.Param("topic")
	payload := ""
	m := echo.Map{}

	if err := c.Bind(&m); err != nil {
		return errors.New(fmt.Sprintf("Error data post: %v", err.Error()))
	}
	payload = utils.EntityToJson(m)

	err := p.useCase.Create(topic, payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"status": "OK"})
}
