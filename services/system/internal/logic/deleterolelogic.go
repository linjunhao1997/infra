package logic

import (
	"context"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/pkg/useraccount"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *v1.DeleteRoleReq) (*v1.CommonResp, error) {
	errx := errorx.ErrRoleDeleteFailed
	roleDao := dao.NewDAO[useraccount.Role](l.svcCtx.DB)
	role, err := roleDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	_, err = roleDao.DeleteByIds(l.ctx, role.Id)
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
