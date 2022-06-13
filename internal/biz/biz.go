package biz

import (
	"context"
	"layout/internal/biz/domain"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase)

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *domain.Greeter) (*domain.Greeter, error)
	Update(context.Context, *domain.Greeter) (*domain.Greeter, error)
	FindByID(context.Context, int64) (*domain.Greeter, error)
	ListByHello(context.Context, string) ([]*domain.Greeter, error)
	ListAll(context.Context) ([]*domain.Greeter, error)
}

type Transaction interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}
