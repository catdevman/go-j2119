package j2119

import (
    "fmt"
    "log"
	"testing"
)

func TestShouldAttachAConditionToAConstraint(t *testing.T) {
    assertion := map[string]string{
       "role":       "R",
       "modal":      "MUST",
       "field_name": "foo",
       "exclude":    "an A, a B, or a C",
    }
    rc := NewRoleConstraints()
    rf := NewRoleFinder()
    matcher := NewMatcher("x")
    af := NewAllowedFields()

    cut := Assigner{
        constraints:   rc,
        roles:         rf,
        matcher:       matcher,
        allowedFields: af,
    }
    for _, v := range []string{"A", "B", "C"}{
        matcher.AddRole(v)
    }
    cut.AssignConstraints(assertion)
    retrieved := rc.Get("R")
    log.Printf("%+v", retrieved)
    //retrieved := cut.constraints.Constraints
    fmt.Println(fmt.Printf("%+v",cut.constraints))
}

func TestNonZeroLessThanConstraint(t *testing.T) {
}

func TestOnlyOneOfConstraint(t *testing.T) {
}

func TestMustHasFieldConstraint(t *testing.T) {
}

func TestMustNotDoesNotHaveFieldConstraint(t *testing.T) {
}

func TestComplexTypeConstraint(t *testing.T) {
}

func TestRecordARelationConstraint(t *testing.T) {
}

func TestRecordAnIsARoleConstraint(t *testing.T) {
}

func TestAssignFieldValueRole(t *testing.T) {
}

func TestChildRoleInAssertion(t *testing.T) {}
