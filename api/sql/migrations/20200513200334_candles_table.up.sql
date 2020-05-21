BEGIN;

CREATE TABLE IF NOT EXISTS candles (
    symbol varchar NOT NULL,
    resolution varchar NOT NULL,
    "from" TIMESTAMP NOT NULL,
    "to" TIMESTAMP NOT NULL,
    "data" JSONB NOT NULL,

    PRIMARY KEY (symbol, resolution, "from", "to")
);

COMMIT;
