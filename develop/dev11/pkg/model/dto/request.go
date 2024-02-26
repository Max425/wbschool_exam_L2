package dto

import (
	"context"
	"net/http"
)

type SuccessClientResponseDto struct {
	Result interface{} `json:"result"`
}

type ErrorClientResponseDto struct {
	Error interface{} `json:"error"`
}

func NewSuccessClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, payload interface{}) {
	sendData(w, statusCode, SuccessClientResponseDto{
		Result: payload,
	})
}

func NewErrorClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, payload interface{}) {
	sendData(w, statusCode, ErrorClientResponseDto{
		Error: payload,
	})
}

func sendData(w http.ResponseWriter, statusCode int, response interface{}) {
	responseJSON, err := response.MarshalJSON()
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
