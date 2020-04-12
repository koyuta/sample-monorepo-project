package handler

import (
	"encoding/json"
	"net/http"
	"path"
)

var Domain = "api.koyuta.com"

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(m string) *ErrorResponse {
	return &ErrorResponse{Message: m}
}

func WriteOKToHeader(rw http.ResponseWriter, resp interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)

	if resp != nil {
		body, _ := json.Marshal(resp)
		rw.Write(body)
	}
}

func WriteCreatedToHeader(rw http.ResponseWriter, loc string) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Location", path.Join(Domain, loc))
	rw.WriteHeader(http.StatusCreated)
}

func WriteBadRequestToHeader(rw http.ResponseWriter, resp *ErrorResponse) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusBadRequest)

	body, _ := json.Marshal(resp)
	rw.Write(body)
}

func WriteNotFoundToHeader(rw http.ResponseWriter, resp *ErrorResponse) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusNotFound)

	body, _ := json.Marshal(resp)
	rw.Write(body)
}

func WriteInternalServerErrorToHeader(rw http.ResponseWriter, resp *ErrorResponse) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusInternalServerError)

	body, _ := json.Marshal(resp)
	rw.Write(body)
}
