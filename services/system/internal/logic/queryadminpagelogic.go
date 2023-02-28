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

type QueryAdminPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAdminPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAdminPageLogic {
	return &QueryAdminPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAdminPageLogic) QueryAdminPage(in *v1.QueryAdminPageReq) (*v1.QueryAdminPageResp, error) {
	errx := errorx.ErrAdminQueryFailed
	adminDao := dao.NewDAO[useraccount.Admin](l.svcCtx.DB)
	page, err := adminDao.FindPage(l.ctx,
		dao.Query{
			Filter: dao.Filter{
				Gte: []dao.Gte{{ColumnName: dao.ColumnCtime, ColumnValue: in.StartCtime}},
				Lte: []dao.Lte{{ColumnName: dao.ColumnCtime, ColumnValue: in.EndCtime}},
			},
			Sorter: &dao.Sorter{
				Columns: []dao.OrderBy{{ColumnName: dao.ColumnCtime, Desc: in.CtimeDesc}},
			},
			Pagination: dao.Pagination{
				CurrentPage: in.CurrentPage,
				PageSize:    in.PageSize,
			},
		}, true)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	data := new(v1.AdminPage)
	err = copier.Copy(data, page)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	return &v1.QueryAdminPageResp{
		Code: 0,
		Msg:  "查询成功",
		Err:  "",
		Data: data,
	}, nil
}
