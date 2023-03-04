package repo

import (
	"context"
	"time"
)

type AccountModel struct {
	ID         int64     `db:"id"`
	TelegramID int64     `db:"telegram_id"`
	Username   string    `db:"username"`
	Fullname   string    `db:"fullname"`
	CreatedAt  time.Time `db:"created_at"`
}

type AccountCreateArgs struct {
	TelegramID int64  `db:"telegram_id"`
	Username   string `db:"username"`
	Fullname   string `db:"fullname"`
}

func (r *Repo) AccountCreate(ctx context.Context, args *AccountCreateArgs) (AccountModel, error) {
	query := `
		INSERT INTO accounts (telegram_id, username, fullname)
		VALUES (:telegram_id, :username, :fullname)
		RETURNING *
	`
	result, err := r.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return AccountModel{}, err
	}

	var account AccountModel
	if err := result.StructScan(&account); err != nil {
		return AccountModel{}, err
	}

	return account, nil
}
