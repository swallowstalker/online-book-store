-- name: CreateUser :one
INSERT INTO "users" ("email", "created_at") VALUES ($1, NOW()) ON CONFLICT(email) DO NOTHING RETURNING *;

-- name: FindUser :one
SELECT * FROM "users" WHERE "email" = $1;