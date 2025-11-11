package handlers

import (
	"net/http"
)

func HandleRoot(sh *SharedContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sh.Stats.IncViewRequests()
		switch r.Method {
		case "GET":
			http.FileServer(http.Dir("public/")).ServeHTTP(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
