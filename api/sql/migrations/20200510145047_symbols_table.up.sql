BEGIN;

CREATE TABLE IF NOT EXISTS symbols (
  id UUID DEFAULT uuid_generate_v4(),

  symbol varchar NOT NULL UNIQUE,
  display_symbol varchar,
  description varchar,

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

  PRIMARY KEY (id)
);

CREATE TRIGGER set_symbols_timestamp BEFORE UPDATE ON symbols FOR EACH ROW EXECUTE PROCEDURE trigger_update_timestamp();

COMMIT;
