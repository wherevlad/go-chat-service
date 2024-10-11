package chat

import (
	"context"

	"github.com/wherevlad/go-chat-service/internal/model"
)

func (s *serv) Get(ctx context.Context, id string) (*model.Chat, error) {
	chat, err := s.chatRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return chat, nil
}
