package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/nicolito128/estacion-meteorologica-server/internal/ucsv"
	"github.com/nicolito128/estacion-meteorologica-server/internal/uhttp"
)

type MeasurementKind string

const (
	// Dato desconocido (no debería ser necesario usarlo)
	UnknownMeasurement MeasurementKind = "unknown"
	// Obtenido por el sensor DHT22
	TempratureMeasurement MeasurementKind = "temperature"
	// Obtenido por el sensor DHT22
	HumidityMeasurement MeasurementKind = "humidity"
	// Obtenido por el pluviómetro
	PrecipitationMeasurement MeasurementKind = "precipitation"
	// Obtenido por el anemómetro
	WindSpeedMeasurement MeasurementKind = "wind_speed"
	// Obtenido por el sensor de presión barométrica
	SeaLevelMeasurement MeasurementKind = "sea_level"
	// Obtenido por el sensor de presión barométrica
	PressureMeasurement MeasurementKind = "pressure"
	// Obtenido por el sensor UV
	UVMeasurement MeasurementKind = "uv"
)

const (
	TemperatureFile   string = "data/temperature.csv"
	HumidityFile      string = "data/humidity.csv"
	PrecipitationFile string = "data/precipitation.csv"
	WindSpeedFile     string = "data/wind_speed.csv"
	SeaLevelFile      string = "data/sea_level.csv"
	PressureFile      string = "data/pressure.csv"
	UVFile            string = "data/uv.csv"
)

// Measurement representa una medición genérica de cualquier tipo.
type Measurement struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
}

type MeasurementsDataResponse struct {
	Kind MeasurementKind `json:"kind"`
	Data []Measurement   `json:"data"`
}

type NewMeasurementDataRequest struct {
	Value float64 `json:"value"`
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

func genericMeasurementHandler(kind MeasurementKind) http.HandlerFunc {
	createFileIfNotExists(kind)
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			data, err := ucsv.ReadAll(getFilepathForMeasurementKind(kind))
			if err != nil {
				http.Error(w, fmt.Errorf("error trying to read temperature file: %v", err).Error(), http.StatusInternalServerError)
				return
			}

			res := MeasurementsDataResponse{
				Kind: kind,
				Data: make([]Measurement, 0),
			}

			for _, line := range data[1:] {
				timestamp := line[0]
				value, err := strconv.ParseFloat(line[2], 64)
				if err != nil {
					http.Error(w, fmt.Errorf("error trying to convert measurement value to float64: %v", err).Error(), http.StatusInternalServerError)
					return
				}
				res.Data = append(res.Data, Measurement{Timestamp: timestamp, Value: value})
			}

			uhttp.WriteJSON(w, http.StatusOK, res)

		case "POST":
			var body NewMeasurementDataRequest

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&body)
			if err != nil {
				http.Error(w, fmt.Errorf("error trying to decode incoming data: %v", err).Error(), http.StatusInternalServerError)
				return
			}

			timestamp := time.Now().Format(time.RFC3339)
			value := strconv.FormatFloat(body.Value, 'f', 4, 64)
			record := []string{timestamp, r.RemoteAddr, value}

			err = ucsv.WriteLine(getFilepathForMeasurementKind(kind), record)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			uhttp.WriteString(w, http.StatusOK, "OK")

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func HandleTemperature() http.HandlerFunc {
	return genericMeasurementHandler(TempratureMeasurement)
}

func HandleHumidity() http.HandlerFunc {
	return genericMeasurementHandler(HumidityMeasurement)
}

func HandlePrecipitation() http.HandlerFunc {
	return genericMeasurementHandler(PrecipitationMeasurement)
}

func HandleWindSpeed() http.HandlerFunc {
	return genericMeasurementHandler(WindSpeedMeasurement)
}

func HandleSeaLevel() http.HandlerFunc {
	return genericMeasurementHandler(SeaLevelMeasurement)
}

func HandlePressure() http.HandlerFunc {
	return genericMeasurementHandler(PressureMeasurement)
}

func HandleUV() http.HandlerFunc {
	return genericMeasurementHandler(UVMeasurement)
}

func getFilepathForMeasurementKind(kind MeasurementKind) string {
	var path string
	switch kind {
	case TempratureMeasurement:
		path = TemperatureFile
	case HumidityMeasurement:
		path = HumidityFile
	case PrecipitationMeasurement:
		path = PrecipitationFile
	case WindSpeedMeasurement:
		path = WindSpeedFile
	case SeaLevelMeasurement:
		path = SeaLevelFile
	case PressureMeasurement:
		path = PressureFile
	case UVMeasurement:
		path = UVFile
	}

	return path
}

func createFileIfNotExists(kind MeasurementKind) error {
	path := getFilepathForMeasurementKind(kind)

	_, err := os.Stat(path)
	if err != nil {
		of, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}

		exfile, err := os.Open(path + ".example")
		if err != nil {
			return fmt.Errorf("error opening example file: %w", err)
		}

		_, err = bufio.NewReader(exfile).WriteTo(of)
		if err != nil {
			return fmt.Errorf("error copying example file: %w", err)
		}
	}

	return nil
}
