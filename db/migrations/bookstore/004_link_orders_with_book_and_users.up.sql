BEGIN;

ALTER TABLE orders ADD CONSTRAINT fk_order_users FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE orders ADD CONSTRAINT fk_order_books FOREIGN KEY (book_id) REFERENCES books(id);

COMMIT;
