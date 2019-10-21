package mycasbin

import (
	"errors"

	"github.com/Ketarin/pg-casbin/app/query"
	cmodel "github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
)

//NewAdapter returns a casbin Adapter
func NewAdapter(manager query.ApiRoleInterface, user string) persist.Adapter {
	return &MyAdapter{
		apiRole: manager,
		user:    user,
	}
}

//MyAdapter represents my casbin adapter
type MyAdapter struct {
	apiRole query.ApiRoleInterface
	user    string
}

//LoadPolicy loads from Manager ( DB or Redis ) all rbac rights
func (a *MyAdapter) LoadPolicy(model cmodel.Model) error {
	//error = not authorized, nothing is really logged here

	//Roles bonds (parent/child)
	globalDefinitions, err := a.apiRole.GetRbacGlobalDefinitions()
	if err != nil {
		return err
	}

	//user bonds ( role attributions )
	userDefinitions, err := a.apiRole.GetRbacUserDefinitions(a.user)
	if err != nil {
		return err
	}

	for _, definition := range append(globalDefinitions, userDefinitions...) {
		persist.LoadPolicyLine(definition, model)
	}

	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *MyAdapter) SavePolicy(model cmodel.Model) error {
	return errors.New("SavePolicy is not implemented")
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *MyAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New("AddPolicy is not implemented")
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *MyAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New("RemovePolicy is not implemented")
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *MyAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("RemoveFilteredPolicy is not implemented")
}
