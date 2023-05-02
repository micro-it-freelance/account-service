package account

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	_ "github.com/jackc/pgx/v5/stdlib"

	core_db "github.com/micro-it-freelance/core/database"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewTestAccountService() Service {
	db := core_db.NewDBConnection()

	return NewAccountService(
		NewAccountRepository(db),
	)
}

func TestAccountService_Create(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := NewTestAccountService()

	table := []struct {
		Name    string
		Account Account
		Handler func(t *testing.T, account Account, err error)
	}{
		{
			Name: "OK",
			Account: Account{
				TelegramID: 1,
				Username:   "ivan",
			},
			Handler: func(t *testing.T, account Account, err error) {
				require.Empty(t, err)
				require.NotEmpty(t, account.ID)
			},
		},
	}

	for _, item := range table {
		t.Run(item.Name, func(t *testing.T) {
			account, err := service.Create(
				ctx,
				item.Account.TelegramID,
				item.Account.Username,
			)
			item.Handler(t, account, err)
		})
	}
}

func TestAccountService_GetByTelegramID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := NewTestAccountService()

	account1, err := service.Create(
		ctx,
		random.Int63n(time.Now().UnixNano()),
		"ivan",
	)
	require.Empty(t, err)

	table := []struct {
		Name    string
		Account Account
		Handler func(t *testing.T, account Account, err error)
	}{
		{
			Name:    "OK",
			Account: account1,
			Handler: func(t *testing.T, account2 Account, err error) {
				require.Empty(t, err)
				require.Equal(t, account1.ID, account2.ID)
				require.Equal(t, account1.TelegramID, account2.TelegramID)
				require.Equal(t, account1.Username, account2.Username)
			},
		},
	}

	for _, item := range table {
		t.Run(item.Name, func(t *testing.T) {
			account, err := service.GetByTelegramID(
				ctx,
				item.Account.TelegramID,
			)
			item.Handler(t, account, err)
		})
	}
}
