package main

import (
	"net/http"
	"sort"

	"github.com/Omorfii/chirpy-boot-project/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {

	var chirps []database.Chirp
	var err error

	query := r.URL.Query().Get("author_id")

	if query != "" {

		id, err := uuid.Parse(query)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "error parsing id")
		}
		chirps, err = cfg.db.GetAllChirpAscFromUserID(r.Context(), id)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "failled to retrieve the chirps")
			return
		}

	} else {

		chirps, err = cfg.db.GetAllChirpAsc(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "failled to retrieve the chirps")
			return
		}
	}

	sortQuerry := r.URL.Query().Get("sort")

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

	if sortQuerry == "desc" {
		sort.Slice(chirpArray, func(i, j int) bool {
			return chirpArray[i].CreatedAt.After(chirpArray[j].CreatedAt)
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
