package main

import "time"

type Device struct {
	Host string
	Data []Measurement
}

func NewDevice(host string) *Device {
	return &Device{Host: host}
}

type Measurement struct {
	Timestamp   time.Time `json:"timestamp"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	WindSpeed   float64   `json:"windSpeed"`
	SeaLevel    float64   `json:"seaLevel"`
	Pressure    float64   `json:"pressure"`
}
