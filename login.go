package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
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

	expire := 3600

	if params.Expires_in_seconds != 0 && params.Expires_in_seconds < 3600 {
		expire = params.Expires_in_seconds
	}

	duration := time.Duration(expire) * time.Second

	token, err := auth.MakeJWT(userDb.ID, cfg.secret, duration)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error trying to make the token")
	}

	user := User{
		ID:        userDb.ID,
		CreatedAt: userDb.CreatedAt,
		UpdatedAt: userDb.UpdatedAt,
		Email:     userDb.Email,
		Token:     token,
	}

	respondWithJSON(w, 200, user)

}
