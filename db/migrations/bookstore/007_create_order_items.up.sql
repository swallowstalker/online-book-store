BEGIN;

CREATE TABLE IF NOT EXISTS order_items (
    "id" BIGSERIAL NOT NULL PRIMARY KEY,
    "order_id" BIGINT NOT NULL,
    "book_id" BIGINT NOT NULL,
    "amount" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_order_books ON order_items(order_id, book_id);

COMMIT;
