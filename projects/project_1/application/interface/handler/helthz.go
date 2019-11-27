package handler

import "net/http"

// Healthz is used by readinessProbe of kubernetes.
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
