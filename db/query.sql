-- name: GetFeed :one
SELECT * FROM feed
WHERE id = ? LIMIT 1;

