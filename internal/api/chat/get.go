package chat

import (
	"context"
	"github.com/wherevlad/go-chat-service/internal/converter"
	desc "github.com/wherevlad/go-chat-service/pkg/chat/v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	chatObj, err := i.chatService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Chat: converter.ToChatFromService(chatObj),
	}, nil
}
