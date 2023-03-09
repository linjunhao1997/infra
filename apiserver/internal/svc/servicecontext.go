package svc

import (
	"infra/apiserver/internal/config"
	"infra/apiserver/internal/middleware"
	"infra/services/system/systemservice"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	SystemSvc           systemservice.SystemService
	AuthorityMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	systemSvc := systemservice.NewSystemService(zrpc.MustNewClient(c.SystemSvc))
	return &ServiceContext{
		Config:              c,
		SystemSvc:           systemSvc,
		AuthorityMiddleware: middleware.NewAuthorityMiddleware(systemSvc).Handle,
	}
}
