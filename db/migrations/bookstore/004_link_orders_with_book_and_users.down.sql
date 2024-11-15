BEGIN;

ALTER TABLE orders DROP CONSTRAINT fk_order_users;
ALTER TABLE orders DROP CONSTRAINT fk_order_books;

COMMIT;
