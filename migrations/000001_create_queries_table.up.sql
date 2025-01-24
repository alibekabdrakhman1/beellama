CREATE TABLE queries (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    response TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
