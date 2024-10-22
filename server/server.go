package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type statusInterceptor struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and calls the original WriteHeader
func (w *statusInterceptor) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func determineColorCode(code int) string {
	var colorCode string = "#777777"

	if strings.HasPrefix(strconv.Itoa(code), "2") {
		colorCode = "#7dfc3d"
	}
	if strings.HasPrefix(strconv.Itoa(code), "3") {
		colorCode = "#3897f5"
	}
	if strings.HasPrefix(strconv.Itoa(code), "4") {
		colorCode = "#ff9c45"
	}
	if strings.HasPrefix(strconv.Itoa(code), "5") {
		colorCode = "#ff4545"
	}
	return colorCode
}

func LogMiddleware(next http.Handler) http.Handler {

	var ipColorCode = "#a0a0a0"
	var ipStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(ipColorCode))

	var pathColorCode = "#3a96dd"
	var pathStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(pathColorCode))

	var splitterStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#606060"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get results of the true handler
		srw := &statusInterceptor{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(srw, r)

		// STATUS
		var s string
		colorCode := determineColorCode(srw.statusCode)
		s += lipgloss.NewStyle().Foreground(lipgloss.Color(colorCode)).Bold(true).Render(strconv.Itoa(srw.statusCode)) + splitterStyle.Render(" | ")

		// IP
		ip := r.Header.Get("CF-Connecting-IP")
		if ip == "" {
			ip = "127.0.0.1"
		}
		s += ipStyle.Render(ip) + splitterStyle.Render(" -> ")

		// PATH
		path := r.URL.Path
		s += pathStyle.Render(path)

		fmt.Print(s + "\n")
	})
}

func main() {
	var fs = http.FileServer(http.Dir("dist"))
	http.Handle("/", LogMiddleware(fs))
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
