package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"project_name/application/controller"
	"project_name/application/interface/handler"
)

// Hello is the api resource for `hello`.
type Hello struct {
	routeContext handler.ContextHandler
	controller   *controller.Hello
}

// NewHello returns new Hello.
func NewHello(r handler.ContextHandler, c *controller.Hello) *Hello {
	return &Hello{routeContext: r, controller: c}
}

// Get replies repeated `hello` string.
func (h *Hello) Get(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()

	var count uint32
	if c, err := strconv.ParseUint(query.Get("count"), 10, 32); err == nil {
		count = uint32(c)
	}

	resp := h.controller.Get(count)

	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
