package utils

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	response := map[string]string{
		"error": err.Error(),
	}

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(data)

	if err != nil {
		ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
