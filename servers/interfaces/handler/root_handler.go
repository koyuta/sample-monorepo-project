package handler

import (
	"net/http"
)

type Root struct{}

func (h *Root) Healthz(rw http.ResponseWriter, r *http.Request) {
	WriteOKToHeader(rw, nil)
}
