package logic

import (
	"context"

	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersInMemberOfGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersInMemberOfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersInMemberOfGroupLogic {
	return &GetUsersInMemberOfGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersInMemberOfGroupLogic) GetUsersInMemberOfGroup(in *user.GetUsersInMemberOfGroupReq) (*user.GetUsersInMemberOfGroupResp, error) {
	ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}

	users, err1 := ldap.GetUsersInMemberOfGroup(in.Group)
	if err1 != nil {
		return nil, err1
	}

	return &user.GetUsersInMemberOfGroupResp{
		Users: users,
	}, nil
}
