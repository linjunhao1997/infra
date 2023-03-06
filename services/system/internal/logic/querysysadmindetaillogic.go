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

type QuerySysadminDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySysadminDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysadminDetailLogic {
	return &QuerySysadminDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuerySysadminDetailLogic) QuerySysadminDetail(in *v1.QuerySysadminDetailReq) (*v1.QuerySysadminDetailResp, error) {
	errx := errorx.ErrSysadminQueryFailed
	adminDao := dao.NewDAO[useraccount.Sysadmin](l.svcCtx.DB)
	admin, err := adminDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	data := &v1.SysadminData{}
	err = copier.Copy(data, admin)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	return &v1.QuerySysadminDetailResp{
		Code: 0,
		Msg:  "查询成功",
		Err:  "",
		Data: data,
	}, nil
}
