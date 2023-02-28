package greeter

import (
	"context"

	"layout/internal/biz/domain"
	"layout/internal/biz/wire"
	"layout/internal/data"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *data.Data
	log  *log.Helper
}

var _ wire.GreeterRepo = (*greeterRepo)(nil)

// NewGreeterRepo .
func NewGreeterRepo(data *data.Data) wire.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(log.GetLogger()),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*domain.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*domain.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*domain.Greeter, error) {
	return nil, nil
}
