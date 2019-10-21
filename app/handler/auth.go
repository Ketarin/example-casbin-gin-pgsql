package handler

import (
	"net/http"

	"github.com/Ketarin/pg-casbin/app/infra"
	"github.com/Ketarin/pg-casbin/app/mycasbin"
	"github.com/Ketarin/pg-casbin/app/query"
	"github.com/gin-gonic/gin"
)

const rbacConf = "/etc/myapp/casbin/rbac_model.conf"

//Auth is a simple handle where you can see everythin
func Auth(conn infra.ConnectionInterface) gin.HandlerFunc {

	return func(c *gin.Context) {

		user := c.GetHeader("AUTH_IDENTITY")

		if user == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, "user is needed")
			return
		}

		// to get DB data in BD
		queryRole := query.NewApiRole(conn)

		// prepare your casbin adapter
		adapter := mycasbin.NewAdapter(queryRole, user)

		// and give it to the enforcer
		enforcer := mycasbin.NewEnforcer(rbacConf, adapter)

		//can you do it ?
		if !enforcer.Enforce(user, c.Request.URL.Path, c.Request.Method) {
			c.AbortWithStatusJSON(http.StatusForbidden, "Forbidden URI")
			return
		}

	}
}
