package errors

import (
	"net/http"
	"volleyapp/domain/models"
)

var BadRequestResponse = models.Response{
	ErrorCode: http.StatusBadRequest,
	Message:   "Bad request",
	Data:      nil,
}

var UnauthorizedResponse = models.Response{
	ErrorCode: http.StatusUnauthorized,
	Message:   "Unauthorized.",
	Data:      nil,
}

var InternalServerErrorResponse = models.Response{
	ErrorCode: http.StatusInternalServerError,
	Message:   "Internal server error.",
	Data:      nil,
}
