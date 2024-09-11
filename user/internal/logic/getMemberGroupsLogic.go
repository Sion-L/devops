package logic

import (
	"context"

	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberGroupsLogic {
	return &GetMemberGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberGroupsLogic) GetMemberGroups(in *user.GetMemberOfGroupsReq) (*user.GetMemberOfGroupsResp, error) {
	ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	groups, err1 := ldap.GetAllMemberOfGroups()
	if err1 != nil {
		return nil, err1
	}

	return &user.GetMemberOfGroupsResp{
		Groups: groups,
	}, nil
}
