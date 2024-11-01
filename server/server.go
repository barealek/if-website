package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/barealek/if-website/chadroom"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type statusInterceptor struct {
	http.ResponseWriter
	statusCode int
}

var chadrooms = make(map[string]*chadroom.Chadroom)

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

func chadgpt(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	// Upgrade to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	var cr = chadroom.NewChadroom()
	cr.Client = ws

	id := uuid.New().String()
	chadrooms[id] = cr

	fmt.Printf("nyt room med id: %v\n", id)

	go func() {
		defer ws.Close()
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			if cr.State != chadroom.StateWaitingUser {
				continue
			}
			cr.ClientSendMessage(string(message))
			cr.SetState(chadroom.StateWaitingChad)
		}
	}()
}

func chadgpthost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	cr, ok := chadrooms[id]
	if !ok {
		http.Error(w, "chadroom not found", http.StatusNotFound)
		return
	}

	cr.Responder = ws
	cr.StartChatting()

	go func() {
		defer ws.Close()
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			if strings.HasPrefix(string(message), "hint::") {
				cr.ClientSendHint(string(message))
				continue
			}

			if cr.State != chadroom.StateWaitingChad {
				continue
			}
			cr.ResponderSendMessage(string(message))
			cr.SetState(chadroom.StateWaitingUser)
		}
	}()
}

func main() {
	var fs = http.FileServer(http.Dir("dist"))
	http.Handle("/", LogMiddleware(fs))
	http.HandleFunc("/api/chad", chadgpt)
	http.HandleFunc("/api/chadhost", chadgpthost)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
