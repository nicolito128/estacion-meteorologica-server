package handlers

import (
	"net/http"
	"time"
)

type MeasurementKind int

const (
	// Dato desconocido (no debería de ser necesario usarlo)
	UnknownMeasurement MeasurementKind = iota

	// Obtenido por el sensor DHT22
	TempratureMeasurement
	// Obtenido por el sensor DHT22
	HumidityMeasurement
	// Obtenido por el pluviómetro
	PrecipitationMeasurement
	// Obtenido por el anemómetro
	WindSpeedMeasurement
	// Obtenido por el sensor de presión barométrica
	SeaLevelMeasurement
	// Obtenido por el sensor de presión barométrica
	PressureMeasurement
	// Obtenido por el sensor UV
	UVMeasurement
)

// Measurement representa una medición genérica de cualquier tipo.
type Measurement struct {
	Timestamp time.Time       `json:"timestamp"`
	Kind      MeasurementKind `json:"kind"`
	Value     float64         `json:"value"`
}

func HandleMeasurements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func HandleTemperature() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleHumidity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandlePrecipitation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleWindSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleSeaLevel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandlePressure() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleUV() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
