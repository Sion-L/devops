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

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetPasswordLogic) ResetPassword(in *user.ResetPasswordReq) (*user.Empty, error) {
	u, err := l.svcCtx.AuthUsersModel.FindOneByUser(l.ctx, in.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, fmt.Errorf("user not found: %s", in.Username)
		}
		return nil, err
	}
	l.Logger.Infof("用户信息: %v", u)
	if u.Source == "ldap" {
		l.Logger.Info("不应该到这")
		return l.resetPasswordInLdap(in.Username, in.OldPassword, in.NewPassword)
	}

	if u.Source == "local" {
		l.Logger.Info("应该到这")
		return l.resetPasswordInMysql(u, in.OldPassword, in.NewPassword)
	}

	return &user.Empty{}, nil
}

func (l *ResetPasswordLogic) resetPasswordInMysql(u *model.AuthUsers, oldPassword, newPassword string) (*user.Empty, error) {
	if u.Password == oldPassword {
		userInfo := &model.AuthUsers{
			UserId:   u.UserId,
			Username: u.Username,
			Password: newPassword,
			NickName: u.NickName,
			RoleName: u.RoleName,
			Source:   u.Source,
			RoleType: u.RoleType,
			Email:    u.Email,
			Mobile:   u.Mobile,
		}
		err := l.svcCtx.AuthUsersModel.Update(l.ctx, userInfo)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("原来的密码不正确")
	}
	return nil, nil
}

func (l *ResetPasswordLogic) resetPasswordInLdap(username, oldPassword, newPassword string) (*user.Empty, error) {
	ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}

	err1 := ldap.ModifyUserPassword(username, oldPassword, newPassword)
	if err1 != nil {
		return nil, err1
	}

	return nil, nil
}
