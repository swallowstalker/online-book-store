-- name: CreateOrder :one
INSERT INTO "orders" ("user_id", "created_at") VALUES ($1, NOW()) RETURNING id, user_id, created_at;

-- name: GetMyOrders :many
SELECT o.id as order_id, o.user_id, u.email as email, o.created_at
FROM "orders" o
JOIN "users" u ON o.user_id = u.id
WHERE o.user_id = $1 ORDER BY o.id DESC LIMIT $2 OFFSET $3;