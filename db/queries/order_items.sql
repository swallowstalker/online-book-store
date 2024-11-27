-- name: CreateOrderItem :one
INSERT INTO "order_items" ("order_id", "book_id", "amount", "created_at") VALUES ($1, $2, $3, NOW()) RETURNING id, order_id, book_id, amount, created_at;
