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
	secret         string
	polka_key      string
}

type errorResponse struct {
	Error string `json:"error"`
}

type parameters struct {
	Body     string    `json:"body"`
	Email    string    `json:"email"`
	User_id  uuid.UUID `json:"user_id"`
	Password string    `json:"password"`
}

type User struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Email         string    `json:"email"`
	Token         string    `json:"token"`
	Refresh_token string    `json:"refresh_token"`
	Is_chirpy_red bool      `json:"is_chirpy_red"`
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	User_id   uuid.UUID `json:"user_id"`
}

type Token struct {
	Token string `json:"token"`
}

type Data struct {
	User_id uuid.UUID `json:"user_id"`
}

type Webhooks struct {
	Event string `json:"event"`
	Data  Data   `json:"data"`
}
