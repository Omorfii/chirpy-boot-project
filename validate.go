package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {

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

	valid := validResponse{
		Valid: true,
	}
	respondWithJSON(w, 200, valid)
}
