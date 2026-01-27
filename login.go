package main

import (
	"encoding/json"
	"net/http"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
)

func (cgf *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Error decoding parameters")
		return
	}

	userDb, err := cgf.db.GetUserFromEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, 401, "Incorrect email or password")
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, userDb.HashedPassword)
	if err != nil || match != true {
		respondWithError(w, 401, "Incorrect email or password")
		return
	}

	user := User{
		ID:        userDb.ID,
		CreatedAt: userDb.CreatedAt,
		UpdatedAt: userDb.UpdatedAt,
		Email:     userDb.Email,
	}

	respondWithJSON(w, 200, user)

}
