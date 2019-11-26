package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/negroni"
)

// PanicFormatter is the recovery formatter for negroni.
type PanicFormatter struct{}

// FormatPanicError returns internal server error response.
func (p *PanicFormatter) FormatPanicError(
	w http.ResponseWriter,
	r *http.Request,
	i *negroni.PanicInformation,
) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("バグった"))
}

// Logging logs error message to stderr.
func Logging(info *negroni.PanicInformation) {
	var err = info.RecoveredPanic
	os.Stderr.WriteString(fmt.Sprint(err))
}
