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

type RemoveUserToMemberOfGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveUserToMemberOfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserToMemberOfGroupLogic {
	return &RemoveUserToMemberOfGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveUserToMemberOfGroupLogic) RemoveUserToMemberOfGroup(req *types.RemoveUserToMemberOfGroupReq) (resp *types.Response, err error) {
	in := &user.RemoveUserToMemberOfGroupReq{
		Username: req.Usernmae,
		Group:    req.Group,
	}

	if _, err = l.svcCtx.User.RemoveUserToMemberOfGroup(l.ctx, in); err != nil {
		return nil, err
	}

	return &types.Response{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("用户%s成功从属性组%s移除", in.Username, in.Group),
	}, nil
}
