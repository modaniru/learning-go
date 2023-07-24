CREATE TABLE users(
    id serial primary key not null,
    login varchar unique not null,
    name varchar not null,
    surname varchar not null
);