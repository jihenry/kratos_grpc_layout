package main

import (
	"layout/internal/biz/service/cron"
	"layout/internal/biz/service/greeter"
	"layout/internal/conf"
	"layout/internal/data"
	rgreeter "layout/internal/data/repo/greeter"
	rserver "layout/internal/data/repo/server"
	"layout/internal/server"

	kafkaConsumer "layout/internal/biz/service/kafka"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
)

//TODO: 目前手动写的，使用wire自动生成
func wireApp(confData *conf.Data, confServer *conf.Server, registry *nacos.Registry) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := rgreeter.NewGreeterRepo(dataData)
	monitorRepo := rserver.NewMonitorRepo(dataData)
	//业务服务
	greeterService := greeter.NewGreeterService(greeterRepo)
	//定时服务
	cronService := cron.NewCronService()
	//kafka服务
	commonKafkaConsumer := kafkaConsumer.NewKafkaCommonConsumer()
	grpcServer := server.NewGRPCServer(confServer, greeterService)
	cronServer, err := server.NewCronServer(cronService)
	if err != nil {
		return nil, nil, err
	}
	monitorServer, err := server.NewMonitorServer(monitorRepo)
	if err != nil {
		return nil, nil, err
	}
	kafkaConsumerServer, err := server.NewKafkaConsumeServer(commonKafkaConsumer)
	if err != nil {
		return nil, nil, err
	}
	app := newApp(registry, grpcServer, cronServer, monitorServer, kafkaConsumerServer)
	return app, func() {
		cleanup()
	}, nil
}
