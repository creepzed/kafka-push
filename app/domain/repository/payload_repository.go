package repository

type PayloadRepository interface {
	Create (topic string, payload string) error
}