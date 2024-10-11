package chat

import (
	"github.com/wherevlad/go-chat-service/internal/client/db"
	"github.com/wherevlad/go-chat-service/internal/repository"
	"github.com/wherevlad/go-chat-service/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

func NewService(
	chatRepository repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
