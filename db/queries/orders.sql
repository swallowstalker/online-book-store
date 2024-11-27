-- name: CreateOrder :one
INSERT INTO "orders" ("user_id", "created_at") VALUES ($1, NOW()) RETURNING id, user_id, created_at;

-- name: GetMyOrders :many
SELECT o.id as order_id, o.user_id, oi.book_id, oi.amount, u.email as email, o.created_at, oi.id as item_id, oi.created_at as item_created_at
FROM "orders" o
JOIN "order_items" oi ON oi.order_id = o.id
JOIN "users" u ON o.user_id = u.id
WHERE o.user_id = $1 LIMIT $2 OFFSET $3;