package errors

import (
	"net/http"
	"volleyapp/internal/core/domain"
)

var BadRequestResponse = &domain.Response{
	ErrorCode: http.StatusBadRequest,
	Message:   "Bad request",
	Data:      nil,
}

var UnauthorizedResponse = domain.Response{
	ErrorCode: http.StatusUnauthorized,
	Message:   "Unauthorized.",
	Data:      nil,
}
