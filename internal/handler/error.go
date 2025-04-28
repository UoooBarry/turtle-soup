package handler

import "net/http"

type (
	InternalHandlerError struct {
		HanlderErrorConfig
	}
	UnauthorizedRequest struct {
		HanlderErrorConfig
	}
	InvalidRequest struct {
		HanlderErrorConfig
	}
	NotfoundError struct {
		HanlderErrorConfig
	}
)

type HanlderErrorConfig struct {
	Message string
	Code    int
}

type HandlerError interface {
	Error() string
	StatusCode() int
}

type CustomHandlerError func(h *HanlderErrorConfig)

func WithCustomMessage(msg string) CustomHandlerError {
	return func(h *HanlderErrorConfig) {
		h.Message = msg
	}
}

func WithCustomCode(code int) CustomHandlerError {
	return func(h *HanlderErrorConfig) {
		h.Code = code
	}
}

func newHanlderErrorConfig(defaultMessage string, defaultCode int, opts ...CustomHandlerError) HanlderErrorConfig {
	he := HanlderErrorConfig{
		Message: defaultMessage, Code: defaultCode,
	}
	for _, opt := range opts {
		opt(&he)
	}
	return he
}

func NewInternalError(opts ...CustomHandlerError) *InternalHandlerError {
	he := newHanlderErrorConfig("internal server error", http.StatusInternalServerError, opts...)

	return &InternalHandlerError{HanlderErrorConfig: he}
}

func NewUnauthorizedError(opts ...CustomHandlerError) *UnauthorizedRequest {
	he := newHanlderErrorConfig("unauthorized", http.StatusUnauthorized, opts...)

	return &UnauthorizedRequest{HanlderErrorConfig: he}
}

func NewNotfoundError(opts ...CustomHandlerError) *NotfoundError {
	he := newHanlderErrorConfig("not found", http.StatusNotFound, opts...)

	return &NotfoundError{HanlderErrorConfig: he}
}

func NewInvalidRequest(opts ...CustomHandlerError) *InvalidRequest {
	he := newHanlderErrorConfig("invalid request", http.StatusBadRequest, opts...)

	return &InvalidRequest{HanlderErrorConfig: he}
}

func (e *HanlderErrorConfig) Error() string {
	return e.Message
}

func (e *HanlderErrorConfig) StatusCode() int {
	return e.Code
}
