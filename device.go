package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type DataEntry struct {
	Timestamp   time.Time `json:"timestamp"`
	Host        string
	Measurement `json:"measurement"`
}

func HandleMeasurement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var data DataEntry
			err := json.NewDecoder(r.Body).Decode(&data.Measurement)
			if err != nil && err != io.EOF {
				http.Error(w, fmt.Errorf("error decoding JSON body: %v", err).Error(), http.StatusBadRequest)
				return
			}
			data.Timestamp = time.Now()
			data.Host = r.RemoteAddr

			file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
			if err != nil {
				http.Error(w, fmt.Errorf("error opening data file: %v", err).Error(), http.StatusInternalServerError)
			}
			defer file.Close()

			wcsv := csv.NewWriter(file)

			wcsv.Write([]string{
				data.Timestamp.Format(time.RFC3339),
				data.Host,
				fmt.Sprintf("%.2f", data.Temperature),
				fmt.Sprintf("%.2f", data.Humidity),
				fmt.Sprintf("%.2f", data.WindSpeed),
				fmt.Sprintf("%.2f", data.SeaLevel),
				fmt.Sprintf("%.2f", data.Pressure),
				fmt.Sprintf("%.2f", data.UV),
			})
			wcsv.Flush()

			fmt.Printf("Received measurement from %s at %s: %+v\n", data.Host, data.Timestamp.Format(time.RFC3339), data.Measurement)
			WriteJSON(w, http.StatusOK, map[string]string{"status": "measurement received"})

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

type Measurement struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	WindSpeed   float64 `json:"windSpeed"`
	SeaLevel    float64 `json:"seaLevel"`
	Pressure    float64 `json:"pressure"`
	UV          float64 `json:"uv"`
}
