package biz

import (
	"context"
	"layout/internal/biz/service/cron"
	"layout/internal/biz/service/greeter"

	"github.com/google/wire"
	"gitlab.yeahka.com/gaas/pkg/mq/kafka"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(greeter.NewGreeterService, kafka.NewKafkaConsumer, cron.NewCronService)

type Transaction interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}
