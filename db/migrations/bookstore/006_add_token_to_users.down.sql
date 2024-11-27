BEGIN;

DROP INDEX IF EXISTS idx_users_token;

ALTER TABLE users DROP COLUMN "password",
    DROP COLUMN "token";

COMMIT;
