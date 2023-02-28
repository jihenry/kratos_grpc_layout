package wire

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/robfig/cron"
)

type CronService interface {
	InitService() error
	Cron(root *cron.Cron) error
}

type GrpcService interface {
	InitService() error
	Register(server *grpc.Server) error
	UnInitService() error //服务结束，用于清理资源
}

type KafkaConsumerService interface {
	OnKafkaMessage(ctx context.Context, pm *sarama.ConsumerMessage) error
}
