package utils

import "net/http"

type ErrResponse struct {
	Status int    `json:"status"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
}

var (
	ErrNotFound = ErrResponse{
		Status: http.StatusNotFound,
		Type:   "Not Found",
	}

	ErrUnauthorized = ErrResponse{
		Status: http.StatusUnauthorized,
		Type:   "Unauthorized",
	}

	ErrForbidden = ErrResponse{
		Status: http.StatusForbidden,
		Type:   "Forbidden",
	}

	ErrBadRequest = ErrResponse{
		Status: http.StatusBadRequest,
		Type:   "Bad Request",
	}

	ErrInternalServer = ErrResponse{
		Status: http.StatusInternalServerError,
		Type:   "Internal Server Error",
	}

	ErrConflict = ErrResponse{
		Status: http.StatusConflict,
		Type:   "Conflict",
	}

	ErrUnprocessable = ErrResponse{
		Status: http.StatusUnprocessableEntity,
		Type:   "Unprocessable Content",
	}
)

func (er *ErrResponse) New(detail string) *ErrResponse {
	return &ErrResponse{
		Status: er.Status,
		Type:   er.Type,
		Detail: detail,
	}
}

func (er *ErrResponse) EchoFormat() (int, ErrResponse) {
	return er.Status, *er
}

func (er *ErrResponse) EchoFormatDetails(detail string) (int, *ErrResponse) {
	er.Detail = detail

	return er.Status, er
}
