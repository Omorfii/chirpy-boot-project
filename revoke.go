package main

import (
	"net/http"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {

	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "error trying to get the token")
		return
	}

	err = cfg.db.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, 401, "token not found")
		return
	}

	w.WriteHeader(204)
}
