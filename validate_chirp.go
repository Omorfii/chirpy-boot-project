package main

import (
	"encoding/json"
	"net/http"

	"github.com/Omorfii/chirpy-boot-project/internal/database"
	"github.com/google/uuid"
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

	id, err := uuid.Parse(params.User_id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error trying to parse the user id")
		return
	}

	parameter := database.CreateChirpParams{
		Body:   cleanedBody,
		UserID: id,
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
