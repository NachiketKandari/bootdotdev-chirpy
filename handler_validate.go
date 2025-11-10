package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedOutput string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	profane := []string{"kerfuffle", "sharbert", "fornax"}

	wordSlice := strings.Split(params.Body, " ")
	for i := 0; i < len(wordSlice); i++ {
		for _, word := range profane {
			if strings.ToLower(wordSlice[i]) == word {
				wordSlice[i] = "****"
				break
			}
		}
	}

	cleanedOutput := strings.Join(wordSlice, " ")

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedOutput: cleanedOutput,
	})
}
