package logic

import (
	"context"

	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserToMemberOfGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserToMemberOfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserToMemberOfGroupLogic {
	return &AddUserToMemberOfGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserToMemberOfGroupLogic) AddUserToMemberOfGroup(in *user.AddUserToMemberOfGroupReq) (*user.Empty, error) {
	ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	err = ldap.AddUserToGroup(in.Username, in.Group)
	if err != nil {
		return nil, err
	}
	l.Logger.Infof("[ldap] success add user %s to %s", in.Username, in.Group)
	return &user.Empty{}, nil
}
