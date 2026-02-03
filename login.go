package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
	"github.com/Omorfii/chirpy-boot-project/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Error decoding parameters")
		return
	}

	userDb, err := cfg.db.GetUserFromEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, 401, "Incorrect email or password")
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, userDb.HashedPassword)
	if err != nil || match != true {
		respondWithError(w, 401, "Incorrect email or password")
		return
	}

	expire := time.Hour
	duration := time.Duration(expire)

	token, err := auth.MakeJWT(userDb.ID, cfg.secret, duration)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error trying to make the token")
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "issue making refresh token")
		return
	}
	expireRefreshToken := time.Now().Add(time.Duration(24) * time.Hour * 60)

	parameter := database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    userDb.ID,
		ExpiredAt: expireRefreshToken,
	}

	refreshTokenDb, err := cfg.db.CreateRefreshToken(r.Context(), parameter)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating refresh token database")
		return
	}

	user := User{
		ID:            userDb.ID,
		CreatedAt:     userDb.CreatedAt,
		UpdatedAt:     userDb.UpdatedAt,
		Email:         userDb.Email,
		Token:         token,
		Refresh_token: refreshTokenDb.Token,
		Is_chirpy_red: userDb.IsChirpyRed,
	}

	respondWithJSON(w, 200, user)

}
