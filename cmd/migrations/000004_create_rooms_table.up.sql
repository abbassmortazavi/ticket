CREATE TABLE rooms
(
    id          SERIAL PRIMARY KEY,
    room_name   VARCHAR(255) NOT NULL,
    description TEXT,
    created_by  INTEGER NOT NULL,
    is_public   BOOLEAN DEFAULT TRUE,
    is_active   BOOLEAN DEFAULT TRUE,
    max_users   INTEGER DEFAULT 50,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);