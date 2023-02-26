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

type QueryRoleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleDetailLogic {
	return &QueryRoleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryRoleDetailLogic) QueryRoleDetail(in *v1.QueryRoleDetailReq) (*v1.QueryRoleDetailResp, error) {
	errx := errorx.ErrRoleQueryFailed
	roleDao := dao.NewDAO[useraccount.Role](l.svcCtx.DB)
	role, err := roleDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	data := &v1.RoleData{}
	err = copier.Copy(data, role)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	return &v1.QueryRoleDetailResp{
		Code: 0,
		Msg:  "查询成功",
		Err:  "",
		Data: data,
	}, nil
}
