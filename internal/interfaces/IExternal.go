package interfaces

import (
	"context"
	"ewallet-framework/external"
)

type IWallet interface {
	CreateWallet(ctx context.Context, UserID int) (*external.Wallet, error)
}
