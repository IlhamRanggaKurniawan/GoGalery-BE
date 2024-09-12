package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, fmt.Sprintf("Error: %s", err.Error()), statusCode)
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
