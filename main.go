package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Zmahl/blog_aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
	fmt.Println(dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	config := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/users", config.handlerUsersCreate)
	mux.HandleFunc("GET /v1/healthz", checkHealth)
	mux.HandleFunc("GET /v1/err", errorResponse)

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	server.ListenAndServe()
}
