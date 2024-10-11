package chat

import (
	"context"
)

func (s *serv) Create(ctx context.Context, usernames []string) (string, error) {
	var id string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.Create(ctx)
		if errTx != nil {
			return errTx
		}

		errTx = s.chatRepository.AddUsersToChat(ctx, id, usernames)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.chatRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return id, nil
}
