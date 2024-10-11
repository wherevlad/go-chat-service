package chat

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/wherevlad/go-chat-service/internal/client/db"
	"github.com/wherevlad/go-chat-service/internal/model"
	"github.com/wherevlad/go-chat-service/internal/repository"
	"github.com/wherevlad/go-chat-service/internal/repository/chat/converter"
	modelRepo "github.com/wherevlad/go-chat-service/internal/repository/chat/model"
)

const (
	chatTableName = "chat"

	chatIDColumn        = "id"
	chatCreatedAtColumn = "created_at"
	chatUpdatedAtColumn = "updated_at"
)

const (
	chatUserTableName = "chat_user"

	chatUserIDColumn        = "id"
	chatUserChatIDColumn    = "chat_id"
	chatUserUsernameColumn  = "username"
	chatUserCreatedAtColumn = "created_at"
	chatUserUpdatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s DEFAULT VALUES RETURNING %s", chatTableName, chatIDColumn)

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id string
	err := r.db.DB().QueryRowContext(ctx, q).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id string) (*model.Chat, error) {
	builder := sq.Select(chatIDColumn, chatCreatedAtColumn, chatUpdatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(chatTableName).
		Where(sq.Eq{chatIDColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository.Get",
		QueryRaw: query,
	}

	var chat modelRepo.Chat
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chat.ID, &chat.CreatedAt, &chat.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToChatFromRepo(&chat), nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	builder := sq.Delete(chatTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errors.Wrap(err, "failed to delete chat")
	}

	return nil
}

func (r *repo) AddUsersToChat(ctx context.Context, chatID string, usernames []string) error {
	builder := sq.Insert(chatUserTableName).
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "username")

	for _, username := range usernames {
		builder = builder.Values(chatID, username)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build chat_user query")
	}

	q := db.Query{
		Name:     "chat_repository.AddUsersToChat",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errors.Wrap(err, "failed to create user chat")
	}

	return nil
}
