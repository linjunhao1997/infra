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

type QueryRolePageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRolePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRolePageLogic {
	return &QueryRolePageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryRolePageLogic) QueryRolePage(in *v1.QueryRolePageReq) (*v1.QueryRolePageResp, error) {
	errx := errorx.ErrRoleQueryFailed
	roleDao := dao.NewDAO[useraccount.Role](l.svcCtx.DB)
	page, err := roleDao.FindPage(l.ctx,
		dao.Query{
			Filter: dao.Filter{
				Gte: []dao.Gte{{ColumnName: dao.ColumnCtime, ColumnValue: in.StartCtime}},
				Lte: []dao.Lte{{ColumnName: dao.ColumnCtime, ColumnValue: in.EndCtime}},
			},
			Sorter: nil,
			Pagination: dao.Pagination{
				CurrentPage: in.CurrentPage,
				PageSize:    in.PageSize,
			},
		}, true)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	data := new(v1.RolePage)
	err = copier.Copy(data, page)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	return &v1.QueryRolePageResp{
		Code: 0,
		Msg:  "查询成功",
		Err:  "",
		Data: data,
	}, nil
}
