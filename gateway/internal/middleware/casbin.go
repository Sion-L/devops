package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
)

// Authorizer stores the casbin handler
type Authorizer struct {
	enforcer  *casbin.Enforcer
	roleField string
}

// AuthorizerOption represents an option.
type AuthorizerOption func(opt *Authorizer)

// WithRoleField returns a custom user unique identity option.
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

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer, opts ...AuthorizerOption) *Authorizer {
	a := &Authorizer{enforcer: e}
	// Initialize Authorizer with provided options
	a.init(opts...)
	return a
}

// GetRole gets the role from the request context.
func (a *Authorizer) GetRole(r *http.Request) (string, bool) {
	role, ok := r.Context().Value(a.roleField).(string)
	return role, ok
}

// CheckPermission checks the role/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
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

// RequirePermission returns the 403 Forbidden to the client.
func (a *Authorizer) RequirePermission(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusForbidden)
	writer.Write([]byte("Forbidden: insufficient permissions"))
}
