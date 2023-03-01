package svc

import (
	"infra/apiserver/internal/config"
	"infra/services/system/systemservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	SystemSvc systemservice.SystemService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		SystemSvc: systemservice.NewSystemService(zrpc.MustNewClient(c.SystemSvc)),
	}
}
