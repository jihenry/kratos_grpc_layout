package biz

import (
	"context"
	"layout/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(log.GetLogger())}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
