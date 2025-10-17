CREATE TABLE IF NOT EXISTS users
(
    id bigserial PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at timestamp(0) with time zone not null DEFAULT NOW(),
    updated_at timestamp(0) with time zone not null DEFAULT NOW()
    );