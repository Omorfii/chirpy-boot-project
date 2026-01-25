package main

import (
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {

	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Resetis omly allow in dev environment"))
		return
	}

	err := cfg.db.DeleteAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to reset database: " + err.Error()))
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	cfg.fileserverHits.Store(0)
	n := cfg.fileserverHits.Load()
	nString := strconv.Itoa(int(n))
	w.Write([]byte("Hits: " + nString))
}
