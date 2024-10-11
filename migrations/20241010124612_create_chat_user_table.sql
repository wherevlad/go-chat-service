-- +goose Up
-- +goose StatementBegin
create table chat_user
(
    id          uuid primary key default gen_random_uuid(),
    chat_id     uuid references chat (id)   on delete cascade on update cascade not null,
    username    varchar(50)                 not null,
    created_at  timestamp                   not null default now(),
    updated_at  timestamp
);
create unique index chat_user_chat_id_author_idx
    on chat_user (chat_id, username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists chat_user;
-- +goose StatementEnd
