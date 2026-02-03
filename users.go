package main

import (
	"encoding/json"
	"net/http"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
	"github.com/Omorfii/chirpy-boot-project/internal/database"
)

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Error decoding parameters")
		return
	}

	hash, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error hashing the password")
		return
	}

	parameters := database.CreateUserParams{
		Email:          params.Email,
		HashedPassword: hash,
	}

	user, err := cfg.db.CreateUser(r.Context(), parameters)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating new user")
		return
	}

	newUser := User{
		ID:            user.ID,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		Email:         user.Email,
		Is_chirpy_red: user.IsChirpyRed,
	}

	respondWithJSON(w, 201, newUser)
}
