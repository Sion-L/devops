package logic

import (
	"context"

	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMemberGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMemberGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMemberGroupLogic {
	return &DelMemberGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMemberGroupLogic) DelMemberGroup(in *user.DelMemberOfGroupReq) (*user.Empty, error) {
	ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	err = ldap.DeleteMemberOfGroup(in.Group)
	if err != nil {
		return nil, err
	}
	l.Logger.Infof("[ldap] success to delete memberof group: %s", in.Group)
	return &user.Empty{}, nil
}
