package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

// ChiContext provides
type ChiContext struct {
	context *chi.Context
}

// NewChiContext returns a new ChiContextHandler.
func NewChiContext() *ChiContext {
	return &ChiContext{context: chi.NewRouteContext()}
}

// URLParam returns
func (c *ChiContext) URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
