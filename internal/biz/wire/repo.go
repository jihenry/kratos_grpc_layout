package wire

import (
	"context"
	"layout/internal/biz/domain"
)

type GreeterRepo interface {
	Save(context.Context, *domain.Greeter) (*domain.Greeter, error)
	Update(context.Context, *domain.Greeter) (*domain.Greeter, error)
	FindByID(context.Context, int64) (*domain.Greeter, error)
	ListByHello(context.Context, string) ([]*domain.Greeter, error)
	ListAll(context.Context) ([]*domain.Greeter, error)
}
