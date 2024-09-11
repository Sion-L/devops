package user

import (
	"context"
	"fmt"
	"github.com/Sion-L/devops/user/user"
	"net/http"

	"github.com/Sion-L/gateway/internal/svc"
	"github.com/Sion-L/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserToMemberOfGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserToMemberOfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserToMemberOfGroupLogic {
	return &AddUserToMemberOfGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserToMemberOfGroupLogic) AddUserToMemberOfGroup(req *types.AddUserToMemberOfGroupReq) (resp *types.Response, err error) {
	in := &user.AddUserToMemberOfGroupReq{
		Username: req.Usernmae,
		Group:    req.Group,
	}

	if _, err = l.svcCtx.User.AddUserToMemberOfGroup(l.ctx, in); err != nil {
		return nil, err
	}

	return &types.Response{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("用户成功填加到%s组", in.Group),
	}, nil
}
