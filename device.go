package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
)

func HandlePing(d *Device) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			data := map[string]uint64{"count": d.PingCount}
			buff, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error trying to marshal data: %v", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", string(buff))
		case "POST":
			atomic.AddUint64(&d.PingCount, 1)
			w.WriteHeader(http.StatusOK)
		}
	}
}

type Device struct {
	PingCount uint64
}

func NewDevice() *Device {
	return &Device{}
}
