/*
Package handlers provides HTTP handlers for the weather station server.
*/
package handlers

import (
	"net/http"

	"github.com/nicolito128/estacion-meteorologica-server/internal/stats"
)

type SharedContext struct {
	Stats *stats.Stats
}

func SetupHandlers(root *http.ServeMux, store *SharedContext) {
	root.HandleFunc("/", HandleRoot(store))

	root.HandleFunc("/stats", HandleStats(store))
	root.HandleFunc("/ping", HandlePing(store))

	root.HandleFunc("/measurements/temperature", HandleTemperature())
	root.HandleFunc("/measurements/humidity", HandleHumidity())
	root.HandleFunc("/measurements/precipitation", HandlePrecipitation())
	root.HandleFunc("/measurements/wind-speed", HandleWindSpeed())
	root.HandleFunc("/measurements/sea-level", HandleSeaLevel())
	root.HandleFunc("/measurements/pressure", HandlePressure())
	root.HandleFunc("/measurements/uv", HandleUV())
}
