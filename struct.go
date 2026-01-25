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

type parameters struct {
	Body string `json:"body"`
}

type cleanedResponse struct {
	Cleaned_body string `json:"cleaned_body"`
}
