package handlers

import "net/http"

type Measurement struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	WindSpeed   float64 `json:"windSpeed"`
	SeaLevel    float64 `json:"seaLevel"`
	Pressure    float64 `json:"pressure"`
	UV          float64 `json:"uv"`
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
