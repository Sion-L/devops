package logic

import (
	"context"

	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMemberGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberGroupLogic {
	return &AddMemberGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMemberGroupLogic) AddMemberGroup(in *user.AddMemberOfGroupReq) (*user.Empty, error) {
	ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	err = ldap.AddNonExistentMemberOfGroup(in.Group)
	if err != nil {
		return nil, err
	}
	l.Logger.Infof("[ldap] success to add memberof %s", in.Group)
	return &user.Empty{}, nil
}
