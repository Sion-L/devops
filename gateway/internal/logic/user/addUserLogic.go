package user

import (
	"context"
	"fmt"
	"github.com/Sion-L/devops/user/user"
	"github.com/Sion-L/gateway/internal/svc"
	"github.com/Sion-L/gateway/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.Response, err error) {
	in := &user.AddUserReq{
		Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Source:   req.Source,
	}
	if _, err = l.svcCtx.User.AddUser(l.ctx, in); err != nil {
		return nil, err
	}

	return &types.Response{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("添加用户%s成功：", in.Username),
	}, nil
}
