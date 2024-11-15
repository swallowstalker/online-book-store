-- name: CreateOrder :one
INSERT INTO "orders" ("user_id", "book_id", "amount", "created_at") VALUES ($1, $2, $3, NOW()) RETURNING *;

-- name: FindOrder :one
SELECT o.*, u.email as email, b.name as book_name
FROM "orders" o
         JOIN "books" b ON o.book_id = b.id
         JOIN "users" u ON o.user_id = u.id
WHERE o.id = $1;

-- name: GetMyOrders :many
SELECT o.*, u.email as email, b.name as book_name
FROM "orders" o
JOIN "books" b ON o.book_id = b.id
JOIN "users" u ON o.user_id = u.id
WHERE o.user_id = $1 LIMIT $2 OFFSET $3;