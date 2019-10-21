package mycasbin

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/persist"
)

//Enforcer reprensents the casbin enforcer
type Enforcer interface {
	Enforce(rvals ...interface{}) bool
}

//NewEnforcer return casbin Enforcer
func NewEnforcer(rbacConf string, adapter persist.Adapter) Enforcer {
	return casbin.NewEnforcer(rbacConf, adapter)
}
