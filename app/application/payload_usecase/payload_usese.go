package payload_usecase

import (
	"github.com/kafka-push/app/domain/repository"
	"github.com/kafka-push/app/shared/log"
)

type PayloadUseCase interface {
	Create(topic, payload string) error
}

type payloadUseCase struct {
	payloadRepository repository.PayloadRepository
}

func NewProductUseCase(payloadRepository repository.PayloadRepository) *payloadUseCase {
	return &payloadUseCase{
		payloadRepository: payloadRepository,
	}
}

func (u *payloadUseCase) Create(topic string, payload string) error {

	log.Info("Adding msg topic %s, value : { %v }", topic, payload)
	err := u.payloadRepository.Create(topic, payload)
	if err != nil {
		return err
	}
	return nil
}
