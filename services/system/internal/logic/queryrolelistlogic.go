package logic

import (
	"context"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/pkg/useraccount"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleListLogic {
	return &QueryRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryRoleListLogic) QueryRoleList(in *v1.QueryRoleListReq) (*v1.QueryRoleListResp, error) {
	errx := errorx.ErrAdminQueryFailed
	roleDao := dao.NewDAO[useraccount.Role](l.svcCtx.DB)
	roles, err := roleDao.Find(l.ctx,
		dao.Filter{
			Gte:  []dao.Gte{{ColumnName: dao.ColumnCtime, ColumnValue: in.StartCtime}},
			Lte:  []dao.Lte{{ColumnName: dao.ColumnCtime, ColumnValue: in.EndCtime}},
			Like: []dao.Like{{ColumnName: "name", ColumnValue: in.Name}},
		},
		nil)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	data := make([]*v1.RoleData, len(roles))
	err = copier.Copy(data, roles)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	return &v1.QueryRoleListResp{
		Code: 0,
		Msg:  "查询成功",
		Err:  "",
		Data: data,
	}, nil
}
