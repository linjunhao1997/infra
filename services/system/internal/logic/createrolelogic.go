package logic

import (
	"context"
	"time"

	"infra/pkg/dao"
	"infra/pkg/errorx"
	"infra/services/system/internal/pkg/useraccount"
	"infra/services/system/internal/store"
	"infra/services/system/internal/svc"
	v1 "infra/services/system/pb/v1"
	"infra/utils"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoleLogic) CreateRole(in *v1.CreateRoleReq) (*v1.CommonIdDataResp, error) {
	errx := errorx.ErrRoleCreateFailed
	if len(in.ForUserType) == 0 {
		return nil, errx.Wrap("角色所属的用户类型必传")
	}
	if len(in.Name) == 0 {
		return nil, errx.Wrap("角色名必传")
	}

	mutex, err := utils.NewRedLockUtil(store.RedisClient).Lock(l.ctx, in.Name+":"+in.ForUserType, 3*time.Second)
	if err != nil {
		return nil, errx.WrapErr(err)
	}
	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			logx.Errorf("unlock failed:%v", err)
		}
	}()

	roleDao := dao.NewDAO[useraccount.Role](l.svcCtx.DB)
	count, err := roleDao.Count(l.ctx, dao.Filter{
		Eq: []dao.Eq{
			{ColumnName: "name", ColumnValue: in.Name},
			{ColumnName: "for_user_type", ColumnValue: in.ForUserType},
		},
	})
	if err != nil {
		return nil, errx.WrapErr(err)
	}

	if count > 0 {
		return nil, errx.Wrap("角色名已存在")
	}

	data := &useraccount.Role{
		Id: dao.GetID(),
	}

	if err := copier.Copy(data, in); err != nil {
		return nil, errx.WrapErr(err)
	}

	if _, err := roleDao.Create(l.ctx, data); err != nil {
		return nil, errx.WrapErr(err)
	}

	return &v1.CommonIdDataResp{
		Code: 0,
		Msg:  "创建成功",
		Err:  "",
		Data: &v1.CommonIdData{
			Id: data.Id,
		},
	}, nil
}
