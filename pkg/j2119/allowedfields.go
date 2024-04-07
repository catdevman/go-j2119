// AllowFields completed
package j2119

type AllowedFields struct {
	allowed map[string][]string
	any     []string
}

func NewAllowedFields() AllowedFields {
	return AllowedFields{}
}

func (af *AllowedFields) SetAllowed(role, child string) {
	if _, ok := af.allowed[role]; !ok {
		af.allowed = make(map[string][]string)
		af.allowed[role] = []string{}
	}
	af.allowed[role] = append(af.allowed[role], child)
}

func (af *AllowedFields) SetAny(role string) {
	af.any = append(af.any, role)
}

func (af *AllowedFields) Allowed(roles []string, child string) bool {
	for _, role := range roles {
		if _, ok := af.allowed[role]; ok && contains(af.allowed[role], child) {
			return true
		}
	}
	return false
}

func (af *AllowedFields) Any(roles []string) bool {
	for _, role := range roles {
		for _, a := range af.any {
			if role == a {
				return true
			}
		}
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
