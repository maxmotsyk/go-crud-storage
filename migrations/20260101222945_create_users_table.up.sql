CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,

    name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,

    age INT NOT NULL CHECK (age >= 0),

    email VARCHAR(255) NOT NULL UNIQUE,

    password TEXT NOT NULL,

    registered_time TIMESTAMPTZ NOT NULL DEFAULT now()
);