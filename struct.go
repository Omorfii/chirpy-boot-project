package main

import (
	"sync/atomic"
	"time"

	"github.com/Omorfii/chirpy-boot-project/internal/database"
	"github.com/google/uuid"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
}

type errorResponse struct {
	Error string `json:"error"`
}

type parameters struct {
	Body  string `json:"body"`
	Email string `json:"email"`
}

type cleanedResponse struct {
	Cleaned_body string `json:"cleaned_body"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}
