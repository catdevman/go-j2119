package j2119

import (
	"fmt"
	"strings"
)

type RoleNotPresentCondition struct {
	ExcludedRoles []string
}

func NewRoleNotPresentCondition(roles []string) RoleNotPresentCondition {
	rpc := RoleNotPresentCondition{}
	rpc.ExcludedRoles = append(rpc.ExcludedRoles, roles...)
	return rpc
}

func (r *RoleNotPresentCondition) String() string {
	return fmt.Sprintf("excluded roles: %s", strings.Join(r.ExcludedRoles, "|"))
}

func (r *RoleNotPresentCondition) ConstraintApplies(node map[string]interface{}, roles []string) bool {
	for _, role := range roles {
		for _, excluded := range r.ExcludedRoles {
			if role == excluded {
				return true
			}
		}
	}
	return false
}
