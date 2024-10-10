-- +goose Up
-- +goose StatementBegin
create table users (
    id bigserial primary key,
    username varchar(50) not null unique,
    password varchar(255) not null,
    name varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
