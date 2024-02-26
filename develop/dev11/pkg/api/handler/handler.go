package handler

import (
	"fmt"
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/constants"
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/service"
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
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", constants.Host)),
	))

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/customers/{uid}/orders", h.customerOrders).Methods("GET")
	apiRouter.HandleFunc("/orders/{uid}", h.order).Methods("GET")
	apiRouter.HandleFunc("/order", h.newOrder).Methods("POST")

	apiRouter.Use(
		h.panicRecoveryMiddleware,
		h.corsMiddleware,
	)

	return r
}
