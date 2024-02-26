package repository

import (
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/model/core"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type EventRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewEventRepository(db *sqlx.DB, log *zap.Logger) *EventRepository {
	return &EventRepository{db: db, log: log}
}

func (r *EventRepository) Create(event *core.Event) (int, error) {
	var id int

	//query := fmt.Sprintf(`INSERT INTO %s (order_uid, data) VALUES ($1, $2) RETURNING id`,
	//	constants.OrderTable)
	//
	//err := r.db.QueryRow(query, event.OrderUID, Event.Data).Scan(&id)
	//if err != nil {
	//	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
	//		r.log.Error(fmt.Sprintf("Event with UID %s already created", Event.OrderUID), zap.Error(err))
	//		return 0, constants.AlreadyExistsError
	//	}
	//	r.log.Error("Error create Event", zap.Error(err))
	//	return 0, err
	//}

	return id, nil
}
