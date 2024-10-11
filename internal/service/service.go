package service

import (
	"context"

	"github.com/wherevlad/go-chat-service/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, usernames []string) (string, error)
	Get(ctx context.Context, id string) (*model.Chat, error)
	Delete(ctx context.Context, id string) error
}
