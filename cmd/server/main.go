package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"layout/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	plog "gitlab.yeahka.com/gaas/pkg/log"
	"gitlab.yeahka.com/gaas/pkg/registry"
	"gitlab.yeahka.com/gaas/pkg/rpc"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(gs *grpc.Server, r *nacos.Registry) *kratos.App {
	rpc.SetDiscovery(r)
	registry.SetDiscovery(r)
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Registrar(r),
		kratos.Server(
			gs,
		),
	)
}

func initLogger(c *conf.Zap) error {
	baseLogger, err := plog.NewZapLogger(
		plog.WithConsole(c.Console),
		plog.WithDir(c.Dir),
		plog.WithFileName(c.FileName),
		plog.WithLevel(c.Level),
	)
	if err != nil {
		return err
	}
	logger := log.With(baseLogger,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	log.SetLogger(logger)
	return nil
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	//1. 加载配置
	bc, err := conf.Load(flagconf)
	defer conf.Close()
	if err != nil {
		panic(err)
	}
	//2. 初始化日志
	if err := initLogger(bc.Zap); err != nil {
		panic(err)
	}
	//3. 注册中心
	registry, err := registry.NewNacosClient(bc.Nacos.Addr, uint64(bc.Nacos.Port),
		constant.WithCacheDir(bc.Nacos.CacheDir),
		constant.WithLogDir(bc.Nacos.LogDir),
		constant.WithNamespaceId(bc.Nacos.Namespace))
	if err != nil {
		panic(err)
	}
	//4. 初始化服务
	app, cleanup, err := wireApp(bc.Server, bc.Data, registry)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	//5. 启动服务
	if err := app.Run(); err != nil {
		panic(err)
	}
}
