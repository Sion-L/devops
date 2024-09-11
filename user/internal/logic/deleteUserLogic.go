package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/model"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserReq) (*user.Empty, error) {
	u, err := l.svcCtx.AuthUsersModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, fmt.Errorf("[mysql] user not found: %v", err)
		}
		return nil, err
	}

	err = l.svcCtx.AuthUsersModel.Delete(l.ctx, in.UserId)
	if err != nil {
		return nil, fmt.Errorf("[mysql] failed to delete user : %v", err)
	}

	if u.Source == "ldap" {
		ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
		if err != nil {
			return nil, err
		}
		err = ldap.DeleteLDAPUser(u.Username)
		if err != nil {
			return nil, err
		}
		l.Logger.Infof("[ldap] success to delete user: %s", u.Username)
	}

	return &user.Empty{}, nil
}
