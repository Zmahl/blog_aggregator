-- name: CreateFeedFollow :one
INSERT INTO feed_follow (id, feed_id, created_at, updated_at, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow
WHERE $1 = id;

-- name: GetFeedFollowsForUser :many
SELECT * FROM feed_follow
WHERE $1 = user_id;