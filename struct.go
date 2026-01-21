package main

import (
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

type errorResponse struct {
	Error string `json:"error"`
}

type validResponse struct {
	Valid bool `json:"valid"`
}

type parameters struct {
	Body string `json:"body"`
}
