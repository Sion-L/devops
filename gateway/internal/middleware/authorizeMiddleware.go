package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type AuthorizeMiddleware struct {
	DataSource string
}

func NewAuthorizeMiddleware(DataSource string) *AuthorizeMiddleware {
	return &AuthorizeMiddleware{
		DataSource: DataSource,
	}
}

func (m *AuthorizeMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mo := model.NewModel()
		mo.AddDef("r", "r", "sub, obj, act")
		mo.AddDef("p", "p", "sub, obj, act")
		mo.AddDef("g", "g", "_, _")
		mo.AddDef("e", "e", "some(where (p.eft == allow))")
		mo.AddDef("m", "m", "g(r.sub, p.sub) && regexMatch(r.obj, p.obj) && regexMatch(r.act, p.act)")

		policy, err := gormadapter.NewAdapter("mysql", m.DataSource, true)
		if err != nil {
			http.Error(w, "Failed to connect to policy database", http.StatusInternalServerError)
			return
		}

		e, err := casbin.NewEnforcer(mo, policy)
		if err != nil {
			http.Error(w, "Failed to create Casbin enforcer", http.StatusInternalServerError)
			return
		}

		a := NewAuthorizer(e, WithRoleField("role"))
		if !a.CheckPermission(r) {
			a.RequirePermission(w)
			return
		}
		next(w, r)
	}
}
