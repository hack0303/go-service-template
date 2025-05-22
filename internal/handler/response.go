package handler

import (
	"encoding/json"
	"net/http"
)

// HTTPResponse represents the standard HTTP response structure
type HTTPResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}) HTTPResponse {
	return HTTPResponse{
		Code: http.StatusOK,
		Msg:  "Success",
		Data: data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code int, msg string, data interface{}) HTTPResponse {
	return HTTPResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// WriteResponse writes the HTTP response in JSON format
func WriteResponse(w http.ResponseWriter, statusCode int, response HTTPResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// ValidateStatusCode validates the status code structure
func ValidateStatusCode(code int) bool {
	if code < 10000000000 || code > 99999999999 {
		return false
	}
	
	// Extract exception category
	category := code / 100000000
	if category != 200 && category != 400 && category != 500 {
		return false
	}
	
	return true
}
