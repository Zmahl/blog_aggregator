-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, api_key, name)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE api_key = $1;

-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;