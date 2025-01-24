-- +goose Up
-- +goose StatementBegin
create table clients
(
    id           bigserial primary key,
    name         varchar(128) not null,
    phone        varchar(12)  not null unique,
    email        varchar(128) not null unique,
    home_address text         not null,
    created_at   timestamptz  not null,
    updated_at   timestamptz  not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table clients;
-- +goose StatementEnd
