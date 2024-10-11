package repository

import (
	"context"

	"github.com/wherevlad/go-chat-service/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context) (string, error)
	Get(ctx context.Context, id string) (*model.Chat, error)
	Delete(ctx context.Context, id string) error
	AddUsersToChat(ctx context.Context, chatID string, usernames []string) error
}
