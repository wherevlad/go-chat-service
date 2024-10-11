package converter

import (
	"github.com/wherevlad/go-chat-service/internal/model"
	desc "github.com/wherevlad/go-chat-service/pkg/chat/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToChatFromService(chat *model.Chat) *desc.Chat {
	var updatedAt *timestamppb.Timestamp
	if chat.UpdatedAt.Valid {
		updatedAt = timestamppb.New(chat.UpdatedAt.Time)
	}

	return &desc.Chat{
		Id:        chat.ID,
		CreatedAt: timestamppb.New(chat.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
