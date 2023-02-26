package main

import (
	"fmt"

	"infra/gateway/internal/config"
	"infra/gateway/internal/handler"
	"infra/gateway/internal/svc"
	consulconfig "infra/pkg/config"
	"infra/pkg/errorx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

func main() {

	var c config.Config
	if err := conf.LoadConfigFromYamlBytes(consulconfig.GetConfigFromConsul(), &c); err != nil {
		panic(err)
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(errorx.GrpcExpectErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
