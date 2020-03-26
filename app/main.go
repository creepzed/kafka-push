package main

import (
	"github.com/kafka-push/app/application/payload_usecase"
	"github.com/kafka-push/app/config"
	"github.com/kafka-push/app/infrastructure/kafka/kafka_repository"
	"github.com/kafka-push/app/interfaces/web"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	kafkaBrokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
)

func main() {
	logFile := config.OpenLogFile()
	defer logFile.Close()

	echoWeb := echo.New()
	echoWeb.HideBanner = true
	echoWeb.Use(middleware.CORS())

	//EndPoints
	web.NewHealthHandler(echoWeb)

	kafkaPayloadRepository := kafka_repository.NewPayloadKafkaRepository(kafkaBrokers...)
	kafkaPayloadUseCase := payload_usecase.NewProductUseCase(kafkaPayloadRepository)
	web.NewPayloadController(echoWeb, kafkaPayloadUseCase)

	log.Println("Starting server")
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	echoWeb.Logger.Fatal(echoWeb.StartServer(server))
}
