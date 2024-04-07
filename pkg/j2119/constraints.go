package j2119

import (
	"fmt"
)

type Condition interface {
	ConstraintApplies(node map[string]interface{}, role string) bool
}

type Constraint struct {
	Conditions []Condition
}

func NewConstraint() *Constraint {
	return &Constraint{
		Conditions: []Condition{},
	}
}

func (c *Constraint) AddCondition(condition Condition) {
	c.Conditions = append(c.Conditions, condition)
}

func (c *Constraint) Applies(node map[string]interface{}, role string) bool {
	if len(c.Conditions) == 0 {
		return true
	}

	for _, condition := range c.Conditions {
		if condition.ConstraintApplies(node, role) {
			return true
		}
	}

	return false
}

type OnlyOneOfConstraint struct {
	*Constraint
	Fields []string
}

func NewOnlyOneOfConstraint(fields []string) *OnlyOneOfConstraint {
	return &OnlyOneOfConstraint{
		Constraint: NewConstraint(),
		Fields:     fields,
	}
}

func (c *OnlyOneOfConstraint) Check(node map[string]interface{}, path string, problems *[]string) {
	count := 0
	for _, field := range c.Fields {
		if _, ok := node[field]; ok {
			count++
		}
	}

	if count > 1 {
		*problems = append(*problems, fmt.Sprintf("%s may have only one of %v", path, c.Fields))
	}
}

type NonEmptyConstraint struct {
	*Constraint
	Name string
}

func NewNonEmptyConstraint(name string) *NonEmptyConstraint {
	return &NonEmptyConstraint{
		Constraint: NewConstraint(),
		Name:       name,
	}
}

func (c *NonEmptyConstraint) Check(node map[string]interface{}, path string, problems *[]string) {
	if value, ok := node[c.Name]; ok {
		if array, ok := value.([]interface{}); ok && len(array) == 0 {
			*problems = append(*problems, fmt.Sprintf("%s.%s is empty, non-empty required", path, c.Name))
		}
	}
}

type HasFieldConstraint struct {
	*Constraint
	Names []string
}

func NewHasFieldConstraint(names []string) *HasFieldConstraint {
	return &HasFieldConstraint{
		Constraint: NewConstraint(),
		Names:      names,
	}
}

func (c *HasFieldConstraint) Check(node map[string]interface{}, path string, problems *[]string) {
	for _, name := range c.Names {
		if _, ok := node[name]; !ok {
			if len(c.Names) == 1 {
				*problems = append(*problems, fmt.Sprintf("%s does not have required field \"%s\"", path, c.Names[0]))
			} else {
				*problems = append(*problems, fmt.Sprintf("%s does not have required field from %v", path, c.Names))
			}
		}
	}
}
