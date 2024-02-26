package dto

import (
	"context"
	"net/http"
)

type RequestInfo struct {
	Status  int
	Message string
}

type SuccessClientResponseDto struct {
	Result interface{} `json:"result"`
}

type ErrorClientResponseDto struct {
	Error interface{} `json:"error"`
}

func NewSuccessClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, payload interface{}) {
	response := SuccessClientResponseDto{
		Result: payload,
	}

	responseJSON, err := response.MarshalJSON()
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	sendData(w, statusCode, responseJSON)
}

func NewErrorClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, payload interface{}) {
	response := ErrorClientResponseDto{
		Error: payload,
	}

	responseJSON, err := response.MarshalJSON()
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	sendData(w, statusCode, responseJSON)
}

func sendData(w http.ResponseWriter, statusCode int, responseJSON []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
