package main

import "net/http"

func errorResponse(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
