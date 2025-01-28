CREATE TABLE messages (
    id VARCHAR(255) PRIMARY KEY,
    content TEXT NOT NULL,
    approved_at TIMESTAMP,
    rejected_at TIMESTAMP
);