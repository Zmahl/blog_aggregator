package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Zmahl/blog_aggregator/internal/database"
	"github.com/Zmahl/blog_aggregator/pkg/feedfetcher"
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

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	config := apiConfig{
		DB: dbQueries,
	}

	worker := feedfetcher.Worker{
		Interval:    2,
		NumberFeeds: 2,
	}

	go worker.FetchAndUpdateFeeds(dbQueries)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/users", config.handlerUsersCreate)
	mux.HandleFunc("POST /v1/feeds", config.middlewareAuth(config.handlerCreateFeed))
	mux.HandleFunc("POST /v1/feed_follows", config.middlewareAuth(config.handlerCreateFeedFollow))

	mux.HandleFunc("GET /v1/feeds", config.handlerGetFeeds)
	mux.HandleFunc("GET /v1/feed_follows", config.middlewareAuth(config.handlerGetFeedFollowsForUser))
	mux.HandleFunc("GET /v1/users", config.middlewareAuth(config.handlerUsersGet))
	mux.HandleFunc("GET /v1/healthz", checkHealth)
	mux.HandleFunc("GET /v1/err", errorResponse)

	mux.HandleFunc("DELETE /v1/feed_follows", config.middlewareAuth(config.handlerDeleteFeedFollow))

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	feedfetcher.FetchDataFromFeed("https://blog.boot.dev/index.xml")
	server.ListenAndServe()
}
