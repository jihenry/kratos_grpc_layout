package server

import (
	"context"
	"layout/internal/conf"
	"layout/internal/server/wire"
	"net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http/pprof"
)

var _ transport.Server = (*grpcServerImpl)(nil)

type grpcServerImpl struct {
	server   *grpc.Server
	services []wire.GrpcService
	pprof    *http.Server
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, services ...wire.GrpcService) transport.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(log.GetLogger()),
		),
		grpc.Timeout(0),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	for _, svc := range services {
		svc.InitService()
	}
	srv := grpc.NewServer(opts...)
	for _, svc := range services {
		svc.Register(srv)
	}
	impl := &grpcServerImpl{
		server:   srv,
		services: services,
	}
	if c.Pprof != nil {
		impl.pprof = &http.Server{
			Addr:    c.Pprof.Addr,
			Handler: pprof.NewHandler(),
		}
	}
	return impl
}

func (s *grpcServerImpl) Start(ctx context.Context) error {
	if s.pprof != nil {
		go s.pprof.ListenAndServe()
	}
	return s.server.Start(ctx)
}

func (s *grpcServerImpl) Stop(ctx context.Context) error {
	if s.pprof != nil {
		s.pprof.Shutdown(ctx)
	}
	if err := s.server.Stop(ctx); err != nil {
		return err
	}
	for _, svc := range s.services {
		svc.UnInitService()
	}
	return nil
}
