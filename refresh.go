package main

import (
	"net/http"
	"time"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {

	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error trying to get the token")
		return
	}

	refreshTokenDb, err := cfg.db.GetRefreshTokenFromTokem(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, 401, "token doesnt exist")
		return
	}
	if refreshTokenDb.ExpiredAt.Before(time.Now()) {
		respondWithError(w, 401, "token is expired")
		return
	}
	if refreshTokenDb.RevokedAt.Valid {
		respondWithError(w, 401, "token is revoked")
		return
	}

	JwtToken, err := auth.MakeJWT(refreshTokenDb.UserID, cfg.secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error trying to make the token")
		return
	}

	token := Token{
		Token: JwtToken,
	}

	respondWithJSON(w, 200, token)
}
