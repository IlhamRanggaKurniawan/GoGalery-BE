package utils

import (
	"net/http"
	"strconv"
)

func GetPathParam(w http.ResponseWriter, r *http.Request, paramsName string, paramsType string) interface{} {
	paramStr := r.PathValue(paramsName)

	if paramStr == "" {
		http.Error(w, "parameter is empty", http.StatusBadRequest)
		return nil
	}

	switch paramsType {
	case "string":
		return paramStr
	case "number":
		paramNum, err := strconv.ParseUint(paramStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid number parameter", http.StatusBadRequest)
			return nil
		}
		return paramNum
	default:
		http.Error(w, "Invalid parameter type", http.StatusBadRequest)
		return nil
	}
}

func GetMultipleQueryParams(w http.ResponseWriter, r *http.Request, params map[string]string) map[string]interface{} {
	results := make(map[string]interface{})
	queryValues := r.URL.Query()

	for paramName, paramType := range params {
		paramStr := queryValues.Get(paramName)

		if paramStr == "" {
			http.Error(w, "Query parameter '"+paramName+"' is empty", http.StatusBadRequest)
			return nil
		}

		switch paramType {
		case "string":
			results[paramName] = paramStr
		case "number":
			paramNum, err := strconv.ParseUint(paramStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid number parameter for '"+paramName+"'", http.StatusBadRequest)
				return nil
			}
			results[paramName] = paramNum
		default:
			http.Error(w, "Invalid parameter type for '"+paramName+"'", http.StatusBadRequest)
			return nil
		}
	}

	return results
}

func GetOneQueryParam(w http.ResponseWriter, r *http.Request, paramName string, paramType string) interface{} {
	queryValues := r.URL.Query()

	paramStr := queryValues.Get(paramName)

	if paramStr == "" {
		http.Error(w, "Query parameter '"+paramName+"' is empty", http.StatusBadRequest)
		return nil
	}

	switch paramType {
	case "string":
		return paramStr
	case "number":
		paramNum, err := strconv.ParseUint(paramStr, 10, 64)

		if err != nil {
			http.Error(w, "Invalid number parameter for '"+paramName+"'", http.StatusBadRequest)
			return nil
		}

		return paramNum
	default:
		http.Error(w, "Invalid parameter type for '"+paramName+"'", http.StatusBadRequest)
		return nil
	}

}
