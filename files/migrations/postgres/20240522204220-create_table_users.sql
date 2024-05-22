
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    users_id bigserial primary key,
    email varchar not null unique,
    username varchar not null unique,
    password varchar null,
    full_name varchar,
    date_of_birth date,
    create_at timestamp(6) default now(),
    update_at timestamp(6) default now()
);
-- +migrate Down
DROP TABLE IF EXISTS users;