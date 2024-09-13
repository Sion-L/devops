package user

import (
	"context"
	"github.com/Sion-L/devops/gateway/internal/svc"
	"github.com/Sion-L/devops/gateway/internal/types"
	"github.com/Sion-L/devops/user/user"
	"github.com/golang-jwt/jwt/v4"
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
	jwtToken, err2 := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.UserId, res.RoleType)
	if err2 != nil {
		return nil, err2
	}
	return &types.LoginResp{
		Username: res.Username,
		JwtToken: types.JwtToken{
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64, roleType int64) (string, error) {

	// 角色映射 塞到token里面去
	roleMap := map[int64]string{
		1: "admin",
		2: "dev",
	}
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["role"] = roleMap[roleType]
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
