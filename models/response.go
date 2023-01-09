package models

import (
	"database/sql"
	"net/http"

	"github.com/lib/pq"
)

// Response ...
type Response struct {
	StatusCode int         `json:"-"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// NewResponse ...
func NewResponse(statusCode int, message string, data interface{}) *Response {
	return &Response{StatusCode: statusCode, Message: message, Data: data}
}

// ParseSQLError ...
func ParseSQLError(err error, message string) *Response {
	code := http.StatusInternalServerError

	if err == sql.ErrNoRows {
		return NewResponse(http.StatusNotFound, "Empty resultset by given query", nil)
	}

	if err, ok := err.(*pq.Error); ok {
		switch err.Code {
		case
			"23000", // integrity_constraint_violation
			"23001", // restrict_violation
			"23502", // not_null_violation
			"23503", // foreign_key_violation
			"23514", // check_violation
			"23P01", // exclusion_violation
			"42501": // insufficient_privilege
			code = http.StatusForbidden
		case "23505": // unique_violation
			code = http.StatusConflict // HTTP Status Conflict
		}
		message = err.Message
	}

	return NewResponse(code, message, nil)
}
