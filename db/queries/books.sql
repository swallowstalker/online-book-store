-- name: GetBooks :many
SELECT * FROM "books" LIMIT $1 OFFSET $2;

-- name: FindBook :one
SELECT * FROM "books" WHERE "id" = $1;