package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and calls the original WriteHeader
func (w *statusResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srw := &statusResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(srw, r)

		var s string
		statusCode := lipgloss.NewStyle().Background(lipgloss.Color("#007000")).Render(strconv.Itoa(srw.statusCode))
		s += statusCode

	})
}

func main() {
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", LogMiddleware(fs))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
