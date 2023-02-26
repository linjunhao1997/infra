package main

import (
	"fmt"

	consulconfig "infra/pkg/config"
	"infra/pkg/errorx"
	"infra/services/system/internal/config"
	"infra/services/system/internal/server"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var c config.Config
	if err := conf.LoadConfigFromYamlBytes(consulconfig.GetConfigFromConsul(), &c); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		v1.RegisterSystemServiceServer(grpcServer, server.NewSystemServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		} else {
			logx.SetLevel(logx.ErrorLevel)
		}
	})
	s.AddUnaryInterceptors(errorx.UnaryExpectErrorInterceptor)
	s.AddStreamInterceptors(errorx.StreamExpectErrorInterceptor)

	defer s.Stop()

	if err := consul.RegisterService(c.ListenOn, c.Consul); err != nil {
		panic(err)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
