package interfaces

import "context"

type ILogoutService interface {
	Logout(ctx context.Context, token string) error
}
