-- +goose Up
-- +goose StatementBegin
create table message
(
    id          uuid primary key default gen_random_uuid(),
    chat_id     uuid references chat (id)   on delete cascade on update cascade not null,
    author      text                        not null,
    content     text                        not null,
    created_at  timestamp                   not null default now(),
    updated_at  timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists message;
-- +goose StatementEnd