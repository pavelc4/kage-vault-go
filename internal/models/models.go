package models

import "time"

type SuccessResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

type ErrorResponse struct {
	Success   bool        `json:"success"`
	Error     ErrorDetail `json:"error"`
	Timestamp time.Time   `json:"timestamp"`
}

type ValidationErrorResponse struct {
	Success   bool            `json:"success"`
	Error     ValidationError `json:"error"`
	Timestamp time.Time       `json:"timestamp"`
}

type ValidationError struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Fields  []FieldError `json:"fields"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
