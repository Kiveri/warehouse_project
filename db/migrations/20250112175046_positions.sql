-- +goose Up
-- +goose StatementBegin
create table positions
(
    id            bigserial primary key,
    name          varchar(128) not null unique,
    barcode       varchar(13)  not null unique,
    price         bigint       not null,
    position_type smallint     not null,
    created_at    timestamptz  not null,
    updated_at    timestamptz  not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS positions
-- +goose StatementEnd
