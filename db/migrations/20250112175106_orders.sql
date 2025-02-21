-- +goose Up
-- +goose StatementBegin
create table orders
(
    id            bigserial primary key,
    positions     jsonb,
    employee_id   bigint,
    client_id     bigint,
    status        smallint,
    delivery_type smallint,
    total         bigint not null,
    created_at    timestamptz not null,
    updated_at    timestamptz not null,

    CONSTRAINT fk_employee
        FOREIGN KEY (employee_id)
            REFERENCES employees (id),

    CONSTRAINT fk_client
        FOREIGN KEY (client_id)
            REFERENCES clients (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table orders;
-- +goose StatementEnd