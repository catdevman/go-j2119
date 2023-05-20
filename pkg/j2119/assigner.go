package j2119

import "fmt"

type Assigner struct {
	constraints   RoleConstraints
	roles         RoleFinder
	matcher       Matcher
	allowedFields AllowedFields
}

func (a *Assigner) New(rc RoleConstraints, rf RoleFinder, m Matcher, af AllowedFields) {
	a.constraints = rc
	a.roles = rf
	a.matcher = m
	a.allowedFields = af
}

func (a *Assigner) AssignRoles(assertion map[string]string) {
	if _, ok := assertion["val_match_present"]; ok {
		// add field value role to
		a.roles.AddFieldValueRole(
			assertion["role"],
			assertion["fieldtomatch"],
			assertion["valtomatch"],
			assertion["newrole"],
		)
		a.matcher.AddRole(assertion["newrole"])
	} else if _, ok := assertion["with_a_field"]; ok {
		a.roles.AddFieldPresenceRole(
			assertion["role"],
			assertion["with_a_field"],
			assertion["newrole"],
		)

		a.matcher.AddRole(assertion["newrole"])
	} else {
		a.roles.AddIsARole(assertion["role"], assertion["newrole"])
		a.matcher.AddRole(assertion["newrole"])
	}
}

func (a *Assigner) AssignConstraints(assertion map[string]string) {
	role := assertion["role"]
	modal := assertion["modal"]
	ty := assertion["type"]
	field_name := assertion["field_name"]
	field_list_string := assertion["field_list"]
	relation := assertion["relation"]
	target := assertion["target"]
	strings := assertion["strings"]
	child_type := assertion["child_type"]
	vals := assertion["vals"]
	fmt.Println(role, modal, ty, field_name, field_list_string, relation, target, strings, child_type, vals)

	var condition any
	if ex, ok := assertion["excluded"]; ok {
		ox := Oxford{}
		excluded_roles := ox.BreakRoleList(a.matcher.roleMatcher, ex)
		condition = RoleNotPresentCondition{}
		c, _ := condition.(RoleNotPresentCondition)
		c.New(excluded_roles)
	}
	// if relation != "" {
        // a.AddRelationConstraint(role, field, relation, target, condition)
    // }

	if role != "" {
	}
}
