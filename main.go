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

func main() {
	flag.Parse()

	if !strings.HasPrefix(*addr, ":") {
		*addr = ":" + *addr
	}

	// Los datos compartidos entre los manejadores
	shared := &handlers.SharedContext{
		Stats: stats.NewStats(),
	}

	// Creamos el servidor principal que servirá nuestras peticiones HTTP
	root := http.NewServeMux()

	// Asignamos los manejadores de las rutas
	handlers.SetupHandlers(root, shared)

	log.Printf("Iniciando servidor en http://127.0.0.1%s/ - CTRL + C para interrumpir", *addr)
	// TODO: Deberíamos manejar TLS en algún momento
	if err := http.ListenAndServe(*addr, root); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
