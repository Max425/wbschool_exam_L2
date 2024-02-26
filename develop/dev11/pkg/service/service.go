package service

import (
	"context"
	"go.uber.org/zap"
)

type Event interface {
	CreateEvent(ctx context.Context, data []byte) (int, error)
	GetEventByUID(ctx context.Context, UID string) (string, error)
	GetCustomerEvents(customerUID string) ([]string, error)
	LoadEventsToCache(ctx context.Context) error
}

type Service struct {
	Event Event
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		Event: NewEventService(repo.Event, repo.Store, log),
	}
}
