package main

import (
	"github.com/kafka-push/app/application/payload_usecase"
	"github.com/kafka-push/app/infrastructure/kafka/kafka_repository"
	"github.com/kafka-push/app/interfaces/web"
	"github.com/kafka-push/app/shared/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	kafkaBrokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
)

func main() {
	echoWeb := echo.New()
	echoWeb.HideBanner = true
	echoWeb.Use(middleware.CORS())

	//EndPoints
	web.NewHealthHandler(echoWeb)

	kafkaPayloadRepository := kafka_repository.NewPayloadKafkaRepository(kafkaBrokers...)
	kafkaPayloadUseCase := payload_usecase.NewProductUseCase(kafkaPayloadRepository)
	web.NewPayloadController(echoWeb, kafkaPayloadUseCase)

	log.Info("Starting Kafka-Push")
	server := &http.Server{
		Addr:         ":8089",
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	echoWeb.Logger.Fatal(echoWeb.StartServer(server))
}
