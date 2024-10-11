-- +goose Up
-- +goose StatementBegin
create table users (
    id bigserial primary key,
    username varchar(50) not null unique,
    password varchar(255) not null,
    name varchar(100) not null,
    role varchar(15) not null,
    avatar varchar(255),
    created_at timestamp not null default (now() at time zone 'UTC'),
    updated_at timestamp not null default (now() at time zone 'UTC'),
    deleted_at timestamp
);

-- fufufafa123
-- $2a$12$lwa8rRCZOvN7neQGECf4n.YSg8AqNNxeFWpY9pyAI9Z2HZgZDYFoi
insert into users (username, password, name, role)
    values ('dev', '$2a$12$lwa8rRCZOvN7neQGECf4n.YSg8AqNNxeFWpY9pyAI9Z2HZgZDYFoi', 'Dev', 'SUPER_ADMIN');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
