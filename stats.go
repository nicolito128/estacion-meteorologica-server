package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

func HandleStats(stats *Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			snaps := stats.Snapshot()
			buff, err := json.Marshal(snaps)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error generating stats: %v", err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", string(buff))
		}
	}
}

type Stats struct {
	StartTime      time.Time
	ViewRequests   uint64
	DeviceRequests uint64
}

func NewStats() *Stats {
	return &Stats{
		StartTime: time.Now(),
	}
}

func (s *Stats) IncViewRequests() {
	atomic.AddUint64(&s.ViewRequests, 1)
}

func (s *Stats) IncDeviceRequests() {
	atomic.AddUint64(&s.DeviceRequests, 1)
}

func (s *Stats) Snapshot() map[string]any {
	return map[string]any{
		"view_requests":   atomic.LoadUint64(&s.ViewRequests),
		"device_requests": atomic.LoadUint64(&s.DeviceRequests),
		"uptime_seconds":  uint64(time.Since(s.StartTime).Seconds()),
	}
}
