package service

import (
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/model/core"
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/repository"
	"go.uber.org/zap"
	"time"
)

type EventService struct {
	repoEvent repository.Event
	log       *zap.Logger
}

func NewEventService(repoEvent repository.Event, log *zap.Logger) *EventService {
	return &EventService{repoEvent: repoEvent, log: log}
}

func (es *EventService) CreateEvent(event *core.Event) (int, error) {
	id, err := es.repoEvent.Create(event)
	if err != nil {
		es.log.Error("Failed to create event", zap.Error(err))
		return 0, err
	}
	return id, nil
}

func (es *EventService) UpdateEvent(event *core.Event) error {
	err := es.repoEvent.Update(event)
	if err != nil {
		es.log.Error("Failed to update event", zap.Error(err))
		return err
	}
	return nil
}

func (es *EventService) DeleteEvent(id int) error {
	err := es.repoEvent.Delete(id)
	if err != nil {
		es.log.Error("Failed to delete event", zap.Error(err))
		return err
	}
	return nil
}

func (es *EventService) GetEventsForTime(date time.Duration) ([]core.Event, error) {
	events, err := es.repoEvent.GetEvents(date)
	if err != nil {
		es.log.Error("Failed to get events for time", zap.Error(err))
		return nil, err
	}
	return events, nil
}
