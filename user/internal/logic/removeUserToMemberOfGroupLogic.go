package logic

import (
	"context"

	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveUserToMemberOfGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveUserToMemberOfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserToMemberOfGroupLogic {
	return &RemoveUserToMemberOfGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveUserToMemberOfGroupLogic) RemoveUserToMemberOfGroup(in *user.RemoveUserToMemberOfGroupReq) (*user.Empty, error) {
	ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	err = ldap.RemoveUserFromGroup(in.Username, in.Group)
	if err != nil {
		return nil, err
	}
	return &user.Empty{}, nil
}
