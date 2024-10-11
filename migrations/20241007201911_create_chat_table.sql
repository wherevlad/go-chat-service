-- +goose Up
-- +goose StatementBegin
create table chat (
    id          uuid primary key default gen_random_uuid(),
    created_at  timestamp not null default now(),
    updated_at  timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists chat;
-- +goose StatementEnd