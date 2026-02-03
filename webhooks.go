package main

import (
	"encoding/json"
	"net/http"

	"github.com/Omorfii/chirpy-boot-project/internal/auth"
)

func (cfg *apiConfig) handlerWebhooks(w http.ResponseWriter, r *http.Request) {

	key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 401, "error getting the apikey")
		return
	}

	if key != cfg.polka_key {
		respondWithError(w, 401, "wrong apikey, access denied")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := Webhooks{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Error decoding parameters")
		return
	}

	if params.Event != "user.upgraded" {
		respondWithError(w, 204, "")
		return
	}

	err = cfg.db.UpgradeUserFromId(r.Context(), params.Data.User_id)
	if err != nil {
		respondWithError(w, 404, "failed to upgrade user to chirp red")
		return
	}

	respondWithJSON(w, 204, "")
}
