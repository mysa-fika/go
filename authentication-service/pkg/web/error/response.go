package response

import (
	"github.com/rs/zerolog"
)

// ErrorResponse structure shared by all handlers
type ErrorResponse struct {
	HTTPStatusCode  int    `json:"-"`
	Status          int    `json:"status"`
	Message         string `json:"message"`
	InternalMessage string `json:"-"`
}

// NewErrorResponse logs the error and returns renderable response.
func NewErrorResponse(msg string, status int, err error, logger *zerolog.Logger) *ErrorResponse {
	internalMessage := ""
	if err != nil {
		logger.Error().Err(err).Msg("HTTP error response")
	}
	return &ErrorResponse{
		HTTPStatusCode:  status,
		Status:          status,
		Message:         msg,
		InternalMessage: internalMessage,
	}
}
