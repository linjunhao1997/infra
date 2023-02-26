package logic

import (
	"context"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"
	"infra/services/system/pkg/useraccount"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAdminLogic {
	return &DeleteAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAdminLogic) DeleteAdmin(in *v1.DeleteAdminReq) (*v1.CommonResp, error) {
	errx := errorx.ErrAdminDeleteFailed
	adminDao := dao.NewDAO[useraccount.Admin](l.svcCtx.DB)
	admin, err := adminDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	_, err = adminDao.DeleteByIds(l.ctx, admin.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	return &v1.CommonResp{
		Code: 0,
		Msg:  "删除成功",
		Err:  "",
		Data: nil,
	}, nil
}
