BEGIN;

CREATE TABLE IF NOT EXISTS login_sessions (
    id UUID DEFAULT uuid_generate_v4(),

    user_id UUID NOT NULL REFERENCES users(id),

    token_id UUID NOT NULL,
    refresh_token varchar NOT NULL,
    active bool NOT NULL DEFAULT true,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id)
);

CREATE TRIGGER set_login_sessions_timestamp BEFORE UPDATE ON login_sessions FOR EACH ROW EXECUTE PROCEDURE trigger_update_timestamp();

COMMIT;
