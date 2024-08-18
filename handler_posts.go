package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Zmahl/blog_aggregator/internal/database"
)

func (cfg *apiConfig) handlerGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Limit int32 `json:"limit"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't request limit")
		return
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  params.Limit,
	})
	fmt.Println(err)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get posts from user")
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
