package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

func HandlePing(d *Device) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			atomic.AddUint64(&d.PingCount, 1)
			w.WriteHeader(http.StatusOK)
			fmt.Println("Ping current count:", d.PingCount)
		}
	}
}

type Device struct {
	PingCount uint64
}

func NewDevice() *Device {
	return &Device{}
}
