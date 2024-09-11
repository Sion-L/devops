package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sion-L/devops/user/model"
	"go.uber.org/zap"

	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserReq) (*user.Empty, error) {
	if in.Source == "ldap" {
		// 不用同步到mysql了 认证源为ldap 用户登录走的是ldap 不过最好调同步接口同步下
		//if err := l.AddUserInMysql(in); err != nil {
		//	return nil, err
		//}
		if err := l.AddUserInLdap(in); err != nil {
			return nil, err
		}
	} else {
		if err := l.AddUserInMysql(in); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (l *AddUserLogic) AddUserInMysql(in *user.AddUserReq) error {
	u := &model.AuthUsers{
		UserId:   0,
		Username: in.Username,
		Password: in.Password,
		NickName: in.NickName,
		RoleName: "管理员",
		Source:   in.Source,
		RoleType: 1,
		Email:    in.Email,
		Mobile:   in.Mobile,
	}
	existingUser, err := l.svcCtx.AuthUsersModel.FindOneByUser(l.ctx, in.Username)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return err
	}

	if existingUser != nil {
		return fmt.Errorf("[mysql] User %s already exists", in.Username)
	} else {
		_, err := l.svcCtx.AuthUsersModel.Insert(l.ctx, u)
		if err != nil {
			return err
		}
		l.Logger.Info("[mysql] Inserted new LDAP user information: ", zap.Any("user", u))
	}

	return nil
}

func (l *AddUserLogic) AddUserInLdap(in *user.AddUserReq) error {
	ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
	if err != nil {
		return err
	}
	err = ldap.AddLDAPUser(in.Username, in.NickName, in.Password, in.Mobile, in.Email)
	if err != nil {
		return err
	}
	return nil
}
