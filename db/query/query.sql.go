// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package query

import (
	"context"
	"database/sql"
)

const getFeed = `-- name: GetFeed :one
SELECT id, title, description, link, feed_link, updated_parsed, image_id, guid FROM feed
WHERE id = ? LIMIT 1
`

func (q *Queries) GetFeed(ctx context.Context, id int64) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeed, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Link,
		&i.FeedLink,
		&i.UpdatedParsed,
		&i.ImageID,
		&i.Guid,
	)
	return i, err
}

const getPerson = `-- name: GetPerson :one
SELECT id, name, email, feed_id FROM person
WHERE id = ? LIMIT 1
`

func (q *Queries) GetPerson(ctx context.Context, id int64) (Person, error) {
	row := q.db.QueryRowContext(ctx, getPerson, id)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.FeedID,
	)
	return i, err
}

const insertFeed = `-- name: InsertFeed :exec
INSERT INTO feed (description, link, feed_link, updated_parsed, guid) values (?,?,?,?,?)
`

type InsertFeedParams struct {
	Description   string
	Link          interface{}
	FeedLink      interface{}
	UpdatedParsed sql.NullString
	Guid          string
}

func (q *Queries) InsertFeed(ctx context.Context, arg InsertFeedParams) error {
	_, err := q.db.ExecContext(ctx, insertFeed,
		arg.Description,
		arg.Link,
		arg.FeedLink,
		arg.UpdatedParsed,
		arg.Guid,
	)
	return err
}
