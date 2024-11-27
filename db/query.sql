-- name: GetFeed :one
SELECT * FROM feed
WHERE id = ? LIMIT 1;


-- name: GetPerson :one
SELECT * FROM person
WHERE id = ? LIMIT 1;

-- name: InsertFeed :exec
INSERT INTO feed (description, link, feed_link, updated_parsed, guid) values (?,?,?,?,?);
