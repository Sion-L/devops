// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "github.com/Sion-L/devops/gateway/internal/handler/user"
	"github.com/Sion-L/devops/gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/addUser",
				Handler: user.AddUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/resetPassword",
				Handler: user.ResetPasswordHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/user/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthorizeMiddleware, serverCtx.RefreshJwtMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/deleteUser",
					Handler: user.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/ldap/addMemberOfGroup",
					Handler: user.AddMemberOfGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/ldap/addSource",
					Handler: user.LdapSourceHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/ldap/addUserToMemberOfGroup",
					Handler: user.AddUserToMemberOfGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/ldap/delMemberOfGroup",
					Handler: user.DelMemberOfGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/ldap/getMemberOfGroups",
					Handler: user.GetMemberOfGroupsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/ldap/getUsersInMemberOfGroup",
					Handler: user.GetUsersInMemberOfGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/ldap/removeUserToMemberOfGroup",
					Handler: user.RemoveUserToMemberOfGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/ldap/verify",
					Handler: user.LdapVerifyHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/user/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.WebSocketMiddleware, serverCtx.AuthorizeMiddleware, serverCtx.RefreshJwtMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/wsTest",
					Handler: user.WsTestHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/user/v1"),
	)
}
