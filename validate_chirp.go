package main

import (
	"encoding/json"
	"net/http"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
	"github.com/Omorfii/chirpy-boot-project/internal/database"
)

func (cfg *apiConfig) handlerValidate(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Error decoding parameters")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	cleanedBody := checkBadWord(params.Body)

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error trying to get the token")
		return
	}
	userId, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, 401, "token is invalid")
		return
	}

	parameter := database.CreateChirpParams{
		Body:   cleanedBody,
		UserID: userId,
	}

	chirpDb, err := cfg.db.CreateChirp(r.Context(), parameter)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error trying to create chirp")
		return
	}

	chirp := Chirp{
		ID:        chirpDb.ID,
		CreatedAt: chirpDb.CreatedAt,
		UpdatedAt: chirpDb.UpdatedAt,
		Body:      chirpDb.Body,
		User_id:   chirpDb.UserID,
	}

	respondWithJSON(w, 201, chirp)
}
