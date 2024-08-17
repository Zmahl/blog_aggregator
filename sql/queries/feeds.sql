-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, last_fetched_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeed :one
SELECT * FROM feeds 
WHERE id = $1;

-- name: GetNextFeedsFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC
LIMIT $1;

-- name: UpdateFeed :exec
UPDATE feeds
SET updated_at = $2, last_fetched_at = $2
WHERE id = $1;