package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetPathParam(r *http.Request, paramName string, paramsType string, errPointer *error) interface{} {
	paramStr := r.PathValue(paramName)

	if paramStr == "" {
		*errPointer = fmt.Errorf("parameter '%s' is empty", paramName)
		return nil
	}

	switch paramsType {
	case "string":
		return paramStr
	case "number":
		paramNum, err := strconv.ParseUint(paramStr, 10, 64)
		if err != nil {
			*errPointer = fmt.Errorf("invalid number parameter for '%s'", paramName)
			return nil
		}
		return paramNum
	default:
		*errPointer = fmt.Errorf("invalid number parameter for '%s'", paramName)
		return nil
	}
}

func GetQueryParam(r *http.Request, paramName string, paramType string, errPointer *error) interface{} {
	queryValues := r.URL.Query()

	paramStr := queryValues.Get(paramName)

	if paramStr == "" {
		*errPointer = fmt.Errorf("query parameter '%s' is empty", paramName)
		return nil
	}

	switch paramType {
	case "string":
		return paramStr
	case "number":
		paramNum, err := strconv.ParseUint(paramStr, 10, 64)
		if err != nil {
			*errPointer = fmt.Errorf("invalid number parameter for '%s'", paramName)
			return nil
		}
		return paramNum
	default:
		*errPointer = fmt.Errorf("invalid parameter type for '%s'", paramName)
		return nil
	}
}
