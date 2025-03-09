package utils

import (
	"encoding/json"
	"net/http"
)

func ParseJSON(w http.ResponseWriter, r *http.Request, v any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(v)
}

func SuccessResponse(v any) map[string]any {
	return map[string]any{
		"success": true,
		"data":    v,
	}
}

func ErrorResponse(err error) map[string]any {
	return map[string]any{
		"success": false,
		"error":   err.Error(),
	}
}
