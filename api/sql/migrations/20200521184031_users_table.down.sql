BEGIN;

DROP TABLE IF EXISTS users;
DROP TRIGGER IF EXISTS set_users_timestamp ON users;
DROP TYPE IF EXISTS user_status;

COMMIT;
