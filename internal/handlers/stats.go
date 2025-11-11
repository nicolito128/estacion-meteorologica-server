package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nicolito128/estacion-meteorologica-server/internal/uhttp"
)

func HandleStats(sh *SharedContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			snaps := sh.Stats.Snapshot()
			_, err := uhttp.WriteJSON(w, http.StatusOK, snaps)
			if err != nil {
				http.Error(w, fmt.Errorf("error trying to marshal stats data: %v", err).Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func HandlePing(sh *SharedContext) http.HandlerFunc {
	stats := sh.Stats
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			data := map[string]int64{"ping": time.Since(stats.StartTime).Milliseconds()}
			_, err := uhttp.WriteJSON(w, http.StatusOK, data)
			if err != nil {
				http.Error(w, fmt.Errorf("error trying to marshal data: %v", err).Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
