-- name: CreateOrder :one
INSERT INTO "orders" ("user_id", "details", "created_at") VALUES ($1, $2, NOW()) RETURNING id, user_id, details, created_at;

-- name: FindOrder :one
SELECT o.*, u.email as email, b.name as book_name
FROM "orders" o
         JOIN "books" b ON o.book_id = b.id
         JOIN "users" u ON o.user_id = u.id
WHERE o.id = $1;

-- name: GetMyOrders :many
SELECT o.*, u.email as email
FROM "orders" o
JOIN "users" u ON o.user_id = u.id
WHERE o.user_id = $1 LIMIT $2 OFFSET $3;