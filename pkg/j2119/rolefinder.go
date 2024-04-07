// barely started, had quite a few questions of the structure of the data for RoleFinder
package j2119

type RoleFinder struct {
	fieldValueRoles    map[string]any
	childRoles         []string // Not sure on this
	grandchildRoles    []string
	fieldPresenceRoles map[string]map[string]string // This inter type might need to be it's own type of a role
	isARoles           map[string][]string
}

func NewRoleFinder() RoleFinder {
	return RoleFinder{}
}

func (f *RoleFinder) AddIsARole(role, otherRole string) {
	if _, exists := f.isARoles[role]; !exists {
		f.isARoles[role] = []string{}
	}
	f.isARoles[role] = append(f.isARoles[role], otherRole)
}

func (f *RoleFinder) AddFieldValueRole(role, field_name, field_value, new_role string) {
	panic("AddFieldValueRole not implemented")
}

func (f *RoleFinder) AddChildRole(role, fieldName, childRole string) {
	panic("AddChildRole not implemented")
}

func (f *RoleFinder) AddGrandchildRole(role, fieldName, childRole string) {
	panic("AddGrandchildRole not implemented")
}

func (f *RoleFinder) FindMoreRoles(node map[string]any, roles *[]string) {
	// Temporary slice to store new roles to avoid modifying the slice during iteration
	newRoles := []string{}

	// Find roles depending on field values
	for _, role := range *roles {
		perFieldName, ok := f.fieldValueRoles[role]
		if ok {
			perFieldNameArr, ok := perFieldName.(map[string]any)
			if ok {
				for fieldName, valueRoles := range perFieldNameArr {
					valueRolesArr, ok := valueRoles.(map[string]string)
					if ok {
						for fieldValue, childRole := range valueRolesArr {
							if val, exists := node[fieldName]; exists && val == fieldValue {
								newRoles = append(newRoles, childRole)
							}
						}
					}
				}
			}
		}
	}

	// Find roles depending on field presence
	for _, role := range *roles {
		perFieldName, ok := f.fieldPresenceRoles[role]
		if ok {
			for fieldName, childRole := range perFieldName {
				if _, exists := node[fieldName]; exists {
					newRoles = append(newRoles, childRole)
				}
			}
		}
	}

	// is_a roles
	for _, role := range *roles {
		otherRoles, ok := f.isARoles[role]
		if ok {
			newRoles = append(newRoles, otherRoles...)
		}
	}

	// Append new roles to the original roles slice
	*roles = append(*roles, newRoles...)
}

func (f *RoleFinder) FindChildRoles(roles []string, fieldName string) []string {
	panic("not implemented")
}

func (f *RoleFinder) FindGrandchildRoles(roles []string, fieldName string) []string {
	panic("not implemented")
}

func (f *RoleFinder) AddFieldPresenceRole(role string, fieldName string, newRole string) {
	panic("not implemented")
}
