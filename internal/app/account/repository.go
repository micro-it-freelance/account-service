package account

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(ctx context.Context, telegramID int64, username string) (account Account, err error)
	GetByTelegramID(ctx context.Context, telegramID int64) (account Account, err error)
}

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) Repository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Create(ctx context.Context, telegramID int64, username string) (account Account, err error) {
	query := `
		INSERT INTO accounts (telegram_id, username) VALUES (:telegram_id, :username) RETURNING *;
	`

	rows, err := r.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"telegram_id": telegramID,
		"username":    username,
	})
	if err != nil || !rows.Next() {
		return Account{}, err
	}

	if err := rows.StructScan(&account); err != nil {
		return Account{}, err
	}

	return account, nil
}

func (r *AccountRepository) GetByTelegramID(ctx context.Context, telegramID int64) (account Account, err error) {
	query := `
		SELECT * FROM accounts WHERE telegram_id=:telegram_id LIMIT 1
	`

	rows, err := r.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"telegram_id": telegramID,
	})
	if err != nil || !rows.Next() {
		return Account{}, err
	}

	if err := rows.StructScan(&account); err != nil {
		return Account{}, err
	}

	return account, nil
}
