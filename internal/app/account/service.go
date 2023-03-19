package account

import "context"

type Service interface {
	Create(ctx context.Context, telegramID int64, username string) (Account, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (account Account, err error)
}

type AccountService struct {
	repo Repository
}

func NewAccountService(r Repository) Service {
	return &AccountService{
		repo: r,
	}
}

func (s *AccountService) Create(ctx context.Context, telegramID int64, username string) (Account, error) {
	return s.repo.Create(ctx, telegramID, username)
}

func (s *AccountService) GetByTelegramID(ctx context.Context, telegramID int64) (account Account, err error) {
	return s.repo.GetByTelegramID(ctx, telegramID)
}
