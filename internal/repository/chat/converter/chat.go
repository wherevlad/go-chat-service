package converter

import (
	"github.com/wherevlad/go-chat-service/internal/model"
	modelRepo "github.com/wherevlad/go-chat-service/internal/repository/chat/model"
)

func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID:        chat.ID,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
}
