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

type MeasurementKind int

const (
	// Dato desconocido (no debería ser necesario usarlo)
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

func (mk MeasurementKind) String() string {
	switch mk {
	case TempratureMeasurement:
		return "temperature"
	case HumidityMeasurement:
		return "humidity"
	case PrecipitationMeasurement:
		return "precipitation"
	case WindSpeedMeasurement:
		return "wind_speed"
	case SeaLevelMeasurement:
		return "sea_level"
	case PressureMeasurement:
		return "pressure"
	case UVMeasurement:
		return "uv"
	default:
		return "unknown"
	}
}

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
	Timestamp string          `json:"timestamp"`
	Kind      MeasurementKind `json:"kind"`
	Value     float64         `json:"value"`
}

type MeasurementsDataResponse struct {
	Data []Measurement `json:"data"`
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

// TODO: convertir el proceso del GET y POST en una función genérica y separada de todo esto
func HandleTemperature() http.HandlerFunc {
	createFileIfNotExists(TempratureMeasurement)
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			data, err := ucsv.ReadAll(TemperatureFile)
			if err != nil {
				http.Error(w, fmt.Errorf("error trying to read temperature file: %v", err).Error(), http.StatusInternalServerError)
				return
			}

			res := MeasurementsDataResponse{
				Data: make([]Measurement, 0),
			}

			for _, line := range data[1:] {
				timestamp := line[0]
				value, err := strconv.ParseFloat(line[2], 64)
				if err != nil {
					http.Error(w, fmt.Errorf("error trying to convert measurement value to float64: %v", err).Error(), http.StatusInternalServerError)
					return
				}
				res.Data = append(res.Data, Measurement{Timestamp: timestamp, Kind: TempratureMeasurement, Value: value})
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

			err = ucsv.WriteLine(TemperatureFile, record)
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

func HandleHumidity() http.HandlerFunc {
	createFileIfNotExists(HumidityMeasurement)
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandlePrecipitation() http.HandlerFunc {
	createFileIfNotExists(PrecipitationMeasurement)
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleWindSpeed() http.HandlerFunc {
	createFileIfNotExists(WindSpeedMeasurement)
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleSeaLevel() http.HandlerFunc {
	createFileIfNotExists(SeaLevelMeasurement)
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandlePressure() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleUV() http.HandlerFunc {
	createFileIfNotExists(UVMeasurement)
	return func(w http.ResponseWriter, r *http.Request) {}
}

func createFileIfNotExists(measurementKind MeasurementKind) error {
	var path string
	switch measurementKind {
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
