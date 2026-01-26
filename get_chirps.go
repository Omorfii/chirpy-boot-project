package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {

	chirps, err := cfg.db.GetAllChirpAsc(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failled to retrieve the chirps")
		return
	}

	var chirpArray []Chirp

	for _, chirp := range chirps {

		chirpArray = append(chirpArray, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			User_id:   chirp.UserID,
		})
	}

	respondWithJSON(w, 200, chirpArray)

}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error parsing chirpID")
		return
	}

	chirp, err := cfg.db.GetChirpFromID(r.Context(), id)
	if err != nil {
		respondWithError(w, 404, "chirp doesnt exist")
		return
	}

	chirpStruct := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		User_id:   chirp.UserID,
	}

	respondWithJSON(w, 200, chirpStruct)
}
