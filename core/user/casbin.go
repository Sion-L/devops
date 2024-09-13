package user

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
)

type Authorizer struct {
	enforcer  *casbin.Enforcer
	roleField string
}

type AuthorizerOption func(opt *Authorizer)

func WithRoleField(roleField string) AuthorizerOption {
	return func(opt *Authorizer) {
		opt.roleField = roleField
	}
}

func (a *Authorizer) init(opts ...AuthorizerOption) {
	a.roleField = "role" // Default role field
	for _, opt := range opts {
		opt(a)
	}
}

func NewAuthorizer(e *casbin.Enforcer, opts ...AuthorizerOption) *Authorizer {
	a := &Authorizer{enforcer: e}
	a.init(opts...)
	return a
}

func (a *Authorizer) GetRole(r *http.Request) (string, bool) {
	role, ok := r.Context().Value(a.roleField).(string)
	return role, ok
}

func (a *Authorizer) CheckPermission(r *http.Request) bool {
	role, ok := a.GetRole(r)
	if !ok {
		return false
	}
	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(role, path, method)
	if err != nil {
		logx.WithContext(r.Context()).Errorf("[CASBIN] enforce error: %s", err.Error())
		return false
	}
	return allowed
}

func (a *Authorizer) RequirePermission(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusForbidden)
	writer.Write([]byte("Forbidden: insufficient permissions"))
}
