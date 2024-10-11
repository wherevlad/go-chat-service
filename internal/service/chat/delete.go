package chat

import (
	"context"
)

func (s *serv) Delete(ctx context.Context, id string) error {
	return s.chatRepository.Delete(ctx, id)
}
