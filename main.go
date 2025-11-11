package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/nicolito128/estacion-meteorologica-server/internal/handlers"
	"github.com/nicolito128/estacion-meteorologica-server/internal/stats"
)

var addr = flag.String("addr", ":8080", "Puerto TCP al que escuchar")

// pass = flag.String("pass", "secret", "Contraseña que debe conocer el dispositivo IoT")

func main() {
	flag.Parse()

	if !strings.HasPrefix(*addr, ":") {
		*addr = ":" + *addr
	}

	shared := &handlers.SharedContext{
		Stats: stats.NewStats(),
	}

	// Creamos el servidor principal que servirá nuestras peticiones HTTP
	root := http.NewServeMux()

	// Asignamos los manejadores de las rutas
	handlers.SetupHandlers(root, shared)

	log.Printf("Iniciando servidor en http://127.0.0.1%s/ - CTRL + C para interrumpir", *addr)
	if err := http.ListenAndServe(*addr, root); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
