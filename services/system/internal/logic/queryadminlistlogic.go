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

type QueryAdminListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAdminListLogic {
	return &QueryAdminListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAdminListLogic) QueryAdminList(in *v1.QueryAdminListReq) (*v1.QueryAdminListResp, error) {
	errx := errorx.ErrAdminQueryFailed
	adminDao := dao.NewDAO[useraccount.Admin](l.svcCtx.DB)
	admins, err := adminDao.Find(l.ctx,
		dao.Filter{
			Gte: []dao.Gte{{ColumnName: dao.ColumnCtime, ColumnValue: in.StartCtime}},
			Lte: []dao.Lte{{ColumnName: dao.ColumnCtime, ColumnValue: in.EndCtime}},
		},
		&dao.Sorter{
			Columns: []dao.OrderBy{{ColumnName: dao.ColumnCtime, Desc: in.CtimeDesc}},
		})
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	data := make([]*v1.AdminData, len(admins))
	err = copier.Copy(data, admins)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	return &v1.QueryAdminListResp{
		Code: 0,
		Msg:  "查询成功",
		Err:  "",
		Data: data,
	}, nil
}
