package greeter

import (
	"context"

	v1 "layout/api/helloworld/v1"
	"layout/internal/biz/domain"
	"layout/internal/biz/wire"
	swire "layout/internal/server/wire"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

var _ swire.GrpcService = (*GreeterService)(nil)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer
	repo wire.GreeterRepo
	log  *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(repo wire.GreeterRepo) *GreeterService {
	return &GreeterService{repo: repo, log: log.NewHelper(log.GetLogger())}
}

// InitService implements wire.GrpcService
func (s *GreeterService) InitService() error {
	return nil
}

// Register implements wire.GrpcService
func (s *GreeterService) Register(server *grpc.Server) error {
	v1.RegisterGreeterServer(server, s)
	return nil
}

// UnInitService implements wire.GrpcService
func (*GreeterService) UnInitService() error {
	return nil
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.repo.Save(ctx, &domain.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
