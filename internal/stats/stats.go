/*
Package stats provides functionality to collect and report server statistics.
*/
package stats

import (
	"sync/atomic"
	"time"
)

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
