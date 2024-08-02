package main

import "net/http"

func checkHealth(w http.ResponseWriter, r *http.Request) {
	type healthResponse struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, healthResponse{
		Status: "ok",
	})
}
