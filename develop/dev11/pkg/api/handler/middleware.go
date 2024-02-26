package handler

import (
	"context"
	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/model/dto"
	"log"
	"net/http"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestInfo := &dto.RequestInfo{}
		childLogger := h.logger.With(zap.String("RequestID", uuid.NewString()))
		ctx := context.WithValue(r.Context(), constants.KeyLogger, childLogger)
		ctx = context.WithValue(ctx, constants.KeyRequestId, uuid.NewString())
		ctx = context.WithValue(ctx, constants.KeyRequestInfo, requestInfo)

		method := r.Method
		path := r.RequestURI
		re := regexp.MustCompile(`\d+`)
		customPath := strings.Split(re.ReplaceAllString(path, "1"), "?")[0]
		h.metrics.Hits.WithLabelValues(customPath, method).Inc()

		next.ServeHTTP(w, r.WithContext(ctx))

		timing := time.Since(start)

		childLogger.Info("Request handled",
			zap.String("Method", method),
			zap.String("RequestURI", path),
			zap.Duration("Time", timing),
		)

		h.metrics.Duration.WithLabelValues(strconv.Itoa(requestInfo.Status), customPath, method).Observe(timing.Seconds())
	})
}

func (h *Handler) panicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger, ok := r.Context().Value(constants.KeyLogger).(*zap.Logger)
				if !ok {
					log.Println("Logger not found in context")
				}

				logger.Error("Panic",
					zap.String("Method", r.Method),
					zap.String("RequestURI", r.RequestURI),
					zap.String("Error", err.(string)),
					zap.String("Message", string(debug.Stack())),
				)
				dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
