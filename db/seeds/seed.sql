BEGIN;

INSERT INTO users (email, created_at)
    VALUES ('pulungragil@gmail.com', NOW()), ('someone1@mail.com', NOW()), ('someone2@mail.com', NOW())
    ON CONFLICT(email) DO NOTHING;

INSERT INTO books (name, created_at)
VALUES ('Chicken Soup of Debugging', NOW()), ('How Google Sheet rules the world', NOW()), ('Catalog of contemporary art', NOW());

COMMIT;