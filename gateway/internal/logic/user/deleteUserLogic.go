package user

import (
	"context"
	"fmt"
	"github.com/Sion-L/devops/user/user"
	"net/http"

	"github.com/Sion-L/devops/gateway/internal/svc"
	"github.com/Sion-L/devops/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserResp) (resp *types.Response, err error) {
	in := &user.DeleteUserReq{UserId: req.UserId}
	if _, err = l.svcCtx.User.DeleteUser(l.ctx, in); err != nil {
		return nil, err
	}
	return &types.Response{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("删除用户%d成功", in.UserId),
	}, nil
}
