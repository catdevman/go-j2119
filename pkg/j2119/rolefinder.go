package j2119

type RoleFinder struct {
	fieldValueRoles    []string // Not sure on this
	childRoles         []string // Not sure on this
	grandchildRoles    []string
	fieldPresenceRoles map[string]map[string]string // This inter type might need to be it's own type of a role
	isARoles           map[string][]string
}

func (f *RoleFinder) AddIsARole(role, otherRole string) {
	if _, ok := f.isARoles[role]; !ok {
		f.isARoles[role] = []string{}
	}
	f.isARoles[role] = append(f.isARoles[role], otherRole)
}

func (f *RoleFinder) AddFieldValueRole(role, field_name, field_value, new_role string) {
	panic("not implemented")
}

func (f *RoleFinder) AddChildRole(role, fieldName, childRole string) {
	panic("not implemented")
}

func (f *RoleFinder) AddGrandchildRole(role, fieldName, childRole string) {
	panic("not implemented")
}

func (f *RoleFinder) FindMoreRoles(node string, roles []string) { // TODO: node is not a string, it will probably be its own type soon
	panic("not implemented")
}

func (f *RoleFinder) FindChildRoles(roles []string, fieldName string) []string {
	panic("not implemented")
	return []string{}
}

func (f *RoleFinder) FindGrandchildRoles(roles []string, fieldName string) []string {
	panic("not implemented")
	return []string{}
}

func (f *RoleFinder) AddFieldPresenceRole(role string, fieldName string, newRole string) {
	if _, ok := f.fieldPresenceRoles[role]; !ok {
		f.fieldPresenceRoles[role] = make(map[string]string)
	}
	f.fieldPresenceRoles[role][fieldName] = newRole
}
