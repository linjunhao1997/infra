package main

import (
	"fmt"
	"net/http"

	"infra/apiserver/internal/config"
	"infra/apiserver/internal/handler"
	"infra/apiserver/internal/svc"
	consulconfig "infra/pkg/config"
	"infra/pkg/errorx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
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

	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Info("global middleware")
			next(w, r)
		}
	})

	httpx.SetErrorHandler(errorx.GrpcExpectErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.DisableStat()
	server.Start()
}
