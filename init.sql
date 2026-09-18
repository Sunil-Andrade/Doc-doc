CREATE TABLE IF NOT EXISTS operations (
    id BIGSERIAL PRIMARY KEY,
    client_id TEXT NOT NULL,
    operation_type TEXT NOT NULL,
    text TEXT NOT NULL,
    sequence BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);