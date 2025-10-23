CREATE TABLE buses
(
    id            SERIAL PRIMARY KEY,
    bus_number    VARCHAR(50) NULL,
    operator_name VARCHAR(100) NULL,
    total_seats   integer NULL,
    bus_type      VARCHAR(100) null,
    amenities     TEXT[],
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);