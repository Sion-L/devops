package user

import (
	"context"
	"github.com/Sion-L/devops/user/user"
	"net/http"

	"github.com/Sion-L/devops/gateway/internal/svc"
	"github.com/Sion-L/devops/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersInMemberOfGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersInMemberOfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersInMemberOfGroupLogic {
	return &GetUsersInMemberOfGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersInMemberOfGroupLogic) GetUsersInMemberOfGroup(req *types.GetUsersInMemberOfGroupReq) (resp *types.ResponseWithData, err error) {
	in := &user.GetUsersInMemberOfGroupReq{
		Group: req.Group,
	}

	groups, err1 := l.svcCtx.User.GetUsersInMemberOfGroup(l.ctx, in)
	if err1 != nil {
		return nil, err1
	}

	return &types.ResponseWithData{
		Code:    http.StatusOK,
		Message: "",
		Data:    groups,
	}, nil
}
