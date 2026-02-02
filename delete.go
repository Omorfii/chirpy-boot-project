package main

import (
	"net/http"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "error trying to get the token")
		return
	}
	userId, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, 401, "token is invalid")
		return
	}

	chirpId, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error parsing chirpID")
		return
	}

	chirp, err := cfg.db.GetChirpFromID(r.Context(), chirpId)
	if err != nil {
		respondWithError(w, 404, "chirp doesnt exist")
		return
	}

	if userId != chirp.UserID {
		respondWithError(w, 403, "not the owner of the chirp")
		return
	}

	err = cfg.db.DeleteChirpFromID(r.Context(), chirp.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error deleting chirp")
		return
	}

	w.WriteHeader(204)

}
