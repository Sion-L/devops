// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/Sion-L/devops/user/internal/logic"
	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) LdapSource(ctx context.Context, in *user.LdapSourceReq) (*user.Empty, error) {
	l := logic.NewLdapSourceLogic(ctx, s.svcCtx)
	return l.LdapSource(in)
}

func (s *UserServer) LdapVerify(ctx context.Context, in *user.LdapVerifyReq) (*user.Empty, error) {
	l := logic.NewLdapVerifyLogic(ctx, s.svcCtx)
	return l.LdapVerify(in)
}

func (s *UserServer) AddUser(ctx context.Context, in *user.AddUserReq) (*user.Empty, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *UserServer) DeleteUser(ctx context.Context, in *user.DeleteUserReq) (*user.Empty, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *UserServer) GetMemberGroups(ctx context.Context, in *user.GetMemberOfGroupsReq) (*user.GetMemberOfGroupsResp, error) {
	l := logic.NewGetMemberGroupsLogic(ctx, s.svcCtx)
	return l.GetMemberGroups(in)
}

func (s *UserServer) AddMemberGroup(ctx context.Context, in *user.AddMemberOfGroupReq) (*user.Empty, error) {
	l := logic.NewAddMemberGroupLogic(ctx, s.svcCtx)
	return l.AddMemberGroup(in)
}

func (s *UserServer) DelMemberGroup(ctx context.Context, in *user.DelMemberOfGroupReq) (*user.Empty, error) {
	l := logic.NewDelMemberGroupLogic(ctx, s.svcCtx)
	return l.DelMemberGroup(in)
}

func (s *UserServer) GetUsersInMemberOfGroup(ctx context.Context, in *user.GetUsersInMemberOfGroupReq) (*user.GetUsersInMemberOfGroupResp, error) {
	l := logic.NewGetUsersInMemberOfGroupLogic(ctx, s.svcCtx)
	return l.GetUsersInMemberOfGroup(in)
}

func (s *UserServer) AddUserToMemberOfGroup(ctx context.Context, in *user.AddUserToMemberOfGroupReq) (*user.Empty, error) {
	l := logic.NewAddUserToMemberOfGroupLogic(ctx, s.svcCtx)
	return l.AddUserToMemberOfGroup(in)
}

func (s *UserServer) RemoveUserToMemberOfGroup(ctx context.Context, in *user.RemoveUserToMemberOfGroupReq) (*user.Empty, error) {
	l := logic.NewRemoveUserToMemberOfGroupLogic(ctx, s.svcCtx)
	return l.RemoveUserToMemberOfGroup(in)
}

func (s *UserServer) ResetPassword(ctx context.Context, in *user.ResetPasswordReq) (*user.Empty, error) {
	l := logic.NewResetPasswordLogic(ctx, s.svcCtx)
	return l.ResetPassword(in)
}
