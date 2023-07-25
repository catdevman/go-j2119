package j2119
import (
    "testing"
)

func TestFailsOnExcludedRoles(t *testing.T) {
    cut := NewRoleNotPresentCondition([]string{"foo", "bar"})
    testJson := map[string]any{
        "bar": "1",
    }
    if ok := cut.ConstraintApplies(testJson, []string{"foo"}); !ok {
        t.Fatal("ConstraintApplies is not correct")
    }

}

func TestSucceedsOnExcludedRoles(t *testing.T) {
    cut := NewRoleNotPresentCondition([]string{"foo", "bar"})
    testJson := map[string]any{
        "bar": "1",
    }
    if ok := cut.ConstraintApplies(testJson, []string{"baz"}); ok {
        t.Fatal("ConstraintApplies function is not correct")
    }

}
