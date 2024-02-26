package repository

import (
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/model/core"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Event interface {
	Create(Event *core.Event) (int, error)
	//GetByUID(UID string) (*core.Event, error)
	//GetCustomerEvents(customerUID string) ([]core.Event, error)
	//GetAll() ([]core.Event, error)
	//DeleteByUID(UID string) error
}

type Repository struct {
	Event
}

func NewRepository(db *sqlx.DB, log *zap.Logger) *Repository {
	return &Repository{
		Event: NewEventRepository(db, log),
	}
}
