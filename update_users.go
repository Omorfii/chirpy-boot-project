package main

import (
	"encoding/json"
	"net/http"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
	"github.com/Omorfii/chirpy-boot-project/internal/database"
)

func (cfg *apiConfig) handlerUpdateUsers(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Error decoding parameters")
		return
	}

	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "error with the access token")
	}
	userId, err := auth.ValidateJWT(tokenString, cfg.secret)
	if err != nil {
		respondWithError(w, 401, "error validating the token")
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error hashing the password")
		return
	}

	parameter := database.UpdateUserInformationParams{
		Email:          params.Email,
		HashedPassword: hashedPassword,
		ID:             userId,
	}

	user, err := cfg.db.UpdateUserInformation(r.Context(), parameter)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating new user")
		return
	}

	updatedUser := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	respondWithJSON(w, 200, updatedUser)

}
