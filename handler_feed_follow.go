package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Zmahl/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't bind feed")
		return
	}

	feed, err := config.DB.GetFeed(r.Context(), params.FeedId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find given feed")
		return
	}

	feedFollow, err := config.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	respondWithJSON(w, http.StatusAccepted, feedFollow)
}

func (config *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		ID uuid.UUID `json:"id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't bind feed follow id")
		return
	}

	err = config.DB.DeleteFeedFollow(r.Context(), params.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't delete the specified feed follow")
	}

	respondWithJSON(w, http.StatusAccepted, "")
}

func (config *apiConfig) handlerGetFeedFollowsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := config.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't get feeds for that user")
		return
	}

	respondWithJSON(w, http.StatusAccepted, feedFollows)
}
