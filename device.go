package main

import "net/http"

func HandlePing(device *Device) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
		}
	}
}

type Device struct {
	Name      string
	PingCount uint64
}

func NewDevice() *Device {
	return &Device{}
}
