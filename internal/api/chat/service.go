package chat

import (
	"github.com/wherevlad/go-chat-service/internal/service"
	desc "github.com/wherevlad/go-chat-service/pkg/chat/v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
