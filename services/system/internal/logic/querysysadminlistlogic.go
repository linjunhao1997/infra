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

type QuerySysadminListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySysadminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysadminListLogic {
	return &QuerySysadminListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuerySysadminListLogic) QuerySysadminList(in *v1.QuerySysadminListReq) (*v1.QuerySysadminListResp, error) {
	errx := errorx.ErrSysadminrQueryFailed
	adminDao := dao.NewDAO[useraccount.Sysadmin](l.svcCtx.DB)
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

	data := make([]*v1.SysadminData, len(admins))
	err = copier.Copy(data, admins)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	return &v1.QuerySysadminListResp{
		Code: 0,
		Msg:  "查询成功",
		Err:  "",
		Data: data,
	}, nil
}
