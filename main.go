package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/nicolito128/estacion-meteorologica-server/internal/handlers"
	"github.com/nicolito128/estacion-meteorologica-server/internal/stats"
)

var addr = flag.String("addr", ":8080", "Puerto TCP al que escuchar")

func main() {
	flag.Parse()

	if !strings.HasPrefix(*addr, ":") {
		*addr = ":" + *addr
	}

	tlsConf := tls.Config{
		MinVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.X25519,
			tls.CurveP256,
		},
	}

	// Los datos compartidos entre los manejadores
	shared := &handlers.SharedContext{
		Stats: stats.NewStats(),
	}

	// Creamos el servidor principal que servirá nuestras peticiones HTTP
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         *addr,
		Handler:      mux,
		TLSConfig:    &tlsConf,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Asignamos los manejadores de las rutas
	handlers.SetupHandlers(mux, shared)

	log.Printf("Iniciando servidor en https://127.0.0.1%s/ - CTRL + C para interrumpir", *addr)
	// TODO: Deberíamos manejar TLS en algún momento
	if err := server.ListenAndServeTLS("certificate.crt", "private.key"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
