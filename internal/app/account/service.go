package account

import "context"

type AccountService struct {
	r *AccountRepository
}

func NewAccountService(r *AccountRepository) *AccountService {
	return &AccountService{
		r: r,
	}
}

func (s *AccountService) Create(ctx context.Context, telegramID int64, username string) error {
	return s.r.Create(ctx, telegramID, username)
}
