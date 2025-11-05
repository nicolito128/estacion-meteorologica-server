package main

import (
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
			_, err := WriteJSON(w, http.StatusOK, snaps)
			if err != nil {
				WriteError(w, http.StatusInternalServerError, fmt.Errorf("error trying to marshal stats data: %v", err))
				return
			}
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

func (s *Stats) TotalRequests() uint64 {
	var total uint64
	atomic.AddUint64(&total, s.DeviceRequests)
	atomic.AddUint64(&total, s.ViewRequests)
	return total
}

func (s *Stats) Snapshot() map[string]any {
	return map[string]any{
		"view_requests":   atomic.LoadUint64(&s.ViewRequests),
		"device_requests": atomic.LoadUint64(&s.DeviceRequests),
		"uptime_seconds":  uint64(time.Since(s.StartTime).Seconds()),
	}
}
