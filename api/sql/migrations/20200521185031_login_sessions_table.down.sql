BEGIN;

DROP TABLE IF EXISTS login_sessions;
DROP TRIGGER IF EXISTS set_login_sessions_timestamp ON login_sessions;

COMMIT;
