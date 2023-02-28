package svc

import (
	"infra/services/system/internal/config"
	"infra/services/system/internal/store"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     store.GormDB,
	}
}

func (ctx *ServiceContext) Service() SystemService {
	return NewSystemService(ctx)
}
