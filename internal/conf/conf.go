package conf

import (
	"gitlab.yeahka.com/gaas/pkg/util"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

var (
	global config.Config
	boot   *Bootstrap
)

func Load(path string) (*Bootstrap, error) {
	global = config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	if err := global.Load(); err != nil {
		return nil, err
	}
	var bc Bootstrap
	if err := global.Scan(&bc); err != nil {
		panic(err)
	}
	boot = &bc
	return boot, nil
}

func Config() config.Config {
	return global
}

func Close() {
	if !util.IsNil(global) {
		global.Close()
	}
}

func Debug() bool {
	return boot.Debug
}

func DataDebug() *Data_Cache {
	if boot.Data.Cache != nil {
		return boot.Data.Cache
	}
	return &Data_Cache{}
}

func Env() Bootstrap_Env {
	return boot.Env
}

func TACfg() *TA {
	return boot.Ta
}
