package account

import (
	"time"

	core_entity "github.com/micro-it-freelance/core/entity"
)

type Account struct {
	ID         int64     `db:"id" json:"db"`
	TelegramID int64     `db:"telegram_id" json:"telegram_id"`
	Username   string    `db:"username" json:"username"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`

	core_entity.ImplementedEntity
}
