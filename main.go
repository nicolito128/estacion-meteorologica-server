package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var addr = flag.String("addr", ":8080", "HTTP network address")

// ename = flag.String("ename", "ESPDevice", "Expected name for ESP devices")
// pass  = flag.String("pass", "secret", "Password for ESP devices")

func main() {
	flag.Parse()

	if !strings.HasPrefix(*addr, ":") {
		*addr = ":" + *addr
	}

	stats := NewStats()
	device := NewDevice()

	http.HandleFunc("/", HandleRoot(stats))
	http.HandleFunc("/stats", HandleStats(stats))
	http.HandleFunc("/ping", HandlePing(device))

	log.Printf("Iniciando servidor en http://127.0.0.1%s/ - CTRL + C para interrumpir", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func HandleRoot(stats *Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats.IncRequests()
		switch r.Method {
		case "GET":
			http.FileServer(http.Dir("public/")).ServeHTTP(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
		}
	}
}
