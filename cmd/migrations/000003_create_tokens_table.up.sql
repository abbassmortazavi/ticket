CREATE TABLE user_tokens
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER     NOT NULL,
    token_type VARCHAR(20) NOT NULL,
    hash_token text        NOT NULL,
    expired_at TIMESTAMP   NOT NULL,
    is_revoked BOOLEAN   DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Indexes for better performance
CREATE INDEX idx_user_tokens_user_id ON user_tokens (user_id);
CREATE INDEX idx_user_tokens_expired_at ON user_tokens (expired_at);
CREATE INDEX idx_user_tokens_hash_token ON user_tokens (hash_token);