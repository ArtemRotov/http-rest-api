-- migrate -path migrations -database "postgres://root:root@localhost/db?sslmode=disable" up
CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);