-- +goose Up
-- +goose StatementBegin
create table employees
(
    id         bigserial primary key,
    name       varchar(128) not null,
    phone      varchar(12)  not null,
    email      varchar(128) not null,
    role       smallint     not null,
    created_at timestamptz  not null,
    updated_at timestamptz  not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table employees;
-- +goose StatementEnd
