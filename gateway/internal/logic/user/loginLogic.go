package user

import (
	"context"
	core "github.com/Sion-L/devops/core/user"
	"github.com/Sion-L/devops/gateway/internal/svc"
	"github.com/Sion-L/devops/gateway/internal/types"
	"github.com/Sion-L/devops/user/user"

	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	in := &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}

	res, err1 := l.svcCtx.User.Login(l.ctx, in)
	if err1 != nil {
		return nil, err1
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	refreshAfter := now + accessExpire/2
	jwtToken, err2 := core.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret,
		now, accessExpire, refreshAfter, res.UserId, res.RoleType)
	if err2 != nil {
		return nil, err2
	}
	return &types.LoginResp{
		Username: res.Username,
		JwtToken: types.JwtToken{
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: refreshAfter,
		},
	}, nil
}
