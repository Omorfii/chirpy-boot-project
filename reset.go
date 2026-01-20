package main

import (
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	cfg.fileserverHits.Store(0)
	n := cfg.fileserverHits.Load()
	nString := strconv.Itoa(int(n))
	w.Write([]byte("Hits: " + nString))
}
