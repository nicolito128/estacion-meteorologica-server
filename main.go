package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	addr = flag.String("addr", ":8080", "Puerto TCP al que escuchar")
	pass = flag.String("pass", "secret", "Contraseña que debe conocer el dispositivo IoT")
)

func main() {
	flag.Parse()

	if !strings.HasPrefix(*addr, ":") {
		*addr = ":" + *addr
	}

	stats := NewStats()

	// Creamos el servidor principal que servirá nuestras peticiones HTTP
	root := http.NewServeMux()

	// Asignamos los manejadores de las rutas
	root.HandleFunc("/", HandleRoot(stats))
	root.HandleFunc("/stats", HandleStats(stats))
	root.HandleFunc("/ping", HandlePing(stats))

	log.Printf("Iniciando servidor en http://127.0.0.1%s/ - CTRL + C para interrumpir", *addr)
	if err := http.ListenAndServe(*addr, root); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func HandleRoot(stats *Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats.IncViewRequests()
		switch r.Method {
		case "GET":
			http.FileServer(http.Dir("public/")).ServeHTTP(w, r)

		default:
			WriteString(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}
}

func HandlePing(stats *Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			data := map[string]int64{"ping": time.Since(stats.StartTime).Milliseconds()}
			_, err := WriteJSON(w, http.StatusOK, data)
			if err != nil {
				WriteError(w, http.StatusInternalServerError, fmt.Errorf("error trying to marshal data: %v", err))
				return
			}
		}
	}
}
