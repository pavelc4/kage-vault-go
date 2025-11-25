package models

type ApiResponse struct {
	Success   bool         `json:"success"`
	Data      interface{}  `json:"data,omitempty"`
	Error     *ErrorDetail `json:"error,omitempty"`
	Timestamp string       `json:"timestamp"`
	Message   string       `json:"message,omitempty"`
}

type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
