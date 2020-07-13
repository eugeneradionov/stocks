BEGIN;

CREATE TYPE user_status AS ENUM ('Active', 'Inactive', 'Blocked', 'Deleted');

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4(),

    email text UNIQUE,
    password text NOT NULL,
    salt varchar NOT NULL,
    name text NOT NULL,

    status user_status NOT NULL DEFAULT 'Inactive',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id)
);

CREATE TRIGGER set_users_timestamp BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE trigger_update_timestamp();

COMMIT;
