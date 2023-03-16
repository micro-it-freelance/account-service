package account

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Create(ctx context.Context, telegramID int64, username string) (id int64, err error) {
	query := `
		INSERT INTO accounts (telegram_id, username) VALUES (:telegram_id, :username) RETURNING id;
	`

	res, err := r.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"telegram_id": telegramID,
		"username":    username,
	})
	if err != nil {
		return 0, err
	}

	if err := res.Scan(&id); err != nil {
		return 0, err
	}
	
	return id, nil
}
