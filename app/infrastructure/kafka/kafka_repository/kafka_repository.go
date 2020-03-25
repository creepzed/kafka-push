package kafka_repository

import (
	"context"
	"fmt"
	"github.com/kafka-push/app/shared/utils"
	"github.com/segmentio/kafka-go"
	_ "github.com/segmentio/kafka-go/snappy"
	log "github.com/sirupsen/logrus"
	"time"
)

type payloadKafkaRepository struct {
	topic   string
	brokers []string
}

func NewPayloadKafkaRepository(brokers ...string) *payloadKafkaRepository {
	return &payloadKafkaRepository{
		brokers: brokers,
	}
}

func (r *payloadKafkaRepository) Create(topic string, payload string) error {
	kafkaMessages := []kafka.Message{
		kafka.Message{
			Key:   []byte(utils.Guid()),
			Value: []byte(payload),
		},
	}
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          r.brokers,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		//CompressionCodec: snappy.NewCompressionCodec(),
		BatchSize:        1,
		BatchTimeout:     10 * time.Millisecond,
	})
	err := kafkaWriter.WriteMessages(context.Background(), kafkaMessages...)
	defer kafkaWriter.Close()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error publishing topic %s and error: %s", topic, err.Error()))
		return err
	}
	return nil
}
