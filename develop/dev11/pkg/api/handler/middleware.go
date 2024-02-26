package handler

import (
	"context"
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/constants"
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/model/dto"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
	"time"
)

func (h *Handler) loggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestInfo := &dto.RequestInfo{}
		ctx := context.WithValue(r.Context(), constants.KeyRequestInfo, requestInfo)

		next.ServeHTTP(w, r.WithContext(ctx))

		timing := time.Since(start)

		requestInfo, ok := ctx.Value(constants.KeyRequestInfo).(*dto.RequestInfo)
		var code int
		if ok {
			code = requestInfo.Status
		}

		h.logger.Info("Request handled",
			zap.Int("StatusCode", code),
			zap.String("RequestURI", r.RequestURI),
			zap.Duration("Time", timing),
		)
	}
}

func (h *Handler) panicRecoveryMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				h.logger.Error("Panic",
					zap.String("Method", r.Method),
					zap.String("RequestURI", r.RequestURI),
					zap.String("Error", err.(string)),
					zap.String("Message", string(debug.Stack())),
				)
				dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
		}()
		next.ServeHTTP(w, r)
	}
}
