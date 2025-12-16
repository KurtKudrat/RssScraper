package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KurtKudrat/RssScraper/internal/db"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
	}

	respondWithJSON(w, 201, user)

}

func (apiCfg *apiConfig) handlerGeteUser(w http.ResponseWriter, r *http.Request, user db.User) {
	// apiKey, err := auth.GetAPIKey(r.Header)
	// if err != nil {
	// 	respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
	// 	return
	// }

	// user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	// if err != nil {
	// 	respondWithError(w, 403, fmt.Sprintf("Couldn't get user: %v", err))
	// 	return
	// }

	respondWithJSON(w, 200, user)
}
