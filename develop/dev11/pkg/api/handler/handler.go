package handler

import (
	"fmt"
	_ "github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/docs"
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/constants"
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/service"
	"github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	services *service.Service
	logger   *zap.Logger
}

func NewHandler(services *service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	r := http.NewServeMux()

	r.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", constants.Host)),
	))

	r.HandleFunc("/create_event", h.Use(h.createEvent))
	r.HandleFunc("/update_event", h.Use(h.updateEvent))
	r.HandleFunc("/delete_event", h.Use(h.deleteEvent))
	r.HandleFunc("/events_for_day", h.Use(h.getEventsForDay))
	r.HandleFunc("/events_for_week", h.Use(h.getEventsForWeek))
	r.HandleFunc("/events_for_month", h.Use(h.getEventsForMonth))

	return r
}

func (h *Handler) Use(next http.HandlerFunc) http.HandlerFunc {
	return h.panicRecoveryMiddleware(
		h.loggingMiddleware(next),
	)
}
