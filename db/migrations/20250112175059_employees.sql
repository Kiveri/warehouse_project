-- +goose Up
-- +goose StatementBegin
create table employees
(
    id         bigserial primary key,
    name       varchar(128) not null,
    phone      varchar(11)  not null unique,
    email      varchar(128) not null unique,
    role       smallint     not null,
    created_at timestamptz  not null,
    updated_at timestamptz  not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table employees;
-- +goose StatementEnd
