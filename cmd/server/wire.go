//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"layout/internal/biz"
	"layout/internal/conf"
	"layout/internal/data"
	"layout/internal/server"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *nacos.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, newApp))
}
