package handler

import "net/http"

// ContextHandler provides
type ContextHandler interface {
	URLParam(*http.Request, string) string
}
