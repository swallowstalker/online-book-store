-- name: GetBooks :many
SELECT * FROM "books" LIMIT $1 OFFSET $2;