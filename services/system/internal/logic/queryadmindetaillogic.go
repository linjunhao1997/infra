package logic

import (
	"context"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"
	"infra/services/system/pkg/useraccount"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAdminDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAdminDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAdminDetailLogic {
	return &QueryAdminDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAdminDetailLogic) QueryAdminDetail(in *v1.QueryAdminDetailReq) (*v1.QueryAdminDetailResp, error) {
	errx := errorx.ErrAdminQueryFailed
	adminDao := dao.NewDAO[useraccount.Admin](l.svcCtx.DB)
	admin, err := adminDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	data := &v1.AdminData{}
	err = copier.Copy(data, admin)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	return &v1.QueryAdminDetailResp{
		Code: 0,
		Msg:  "查询成功",
		Err:  "",
		Data: data,
	}, nil
}
