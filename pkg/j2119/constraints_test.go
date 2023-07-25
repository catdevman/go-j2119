package j2119

import (
	"testing"
)

func TestConstraint(t *testing.T) {
	// cut := NewHasFieldConstraint("foo")
	// var jsonRaw map[string]interface{}
	// json.Unmarshal([]byte(`{ "bar": 1 }`), &jsonRaw)
	//
	// if !cut.Applies(jsonRaw, "foo") {
	// 	t.Errorf("Expected Applies to return true, got false")
	// }
	//
	// cond := NewRoleNotPresentCondition([]string{"foo", "bar"})
	// cut.AddCondition(cond)
	//
	// if cut.Applies(jsonRaw, []string{"foo"}) {
	// 	t.Errorf("Expected Applies to return false, got true")
	// }
	//
	// if !cut.Applies(jsonRaw, []string{"baz"}) {
	// 	t.Errorf("Expected Applies to return true, got false")
	// }
}

func TestHasFieldConstraint(t *testing.T) {
	// cut := NewHasFieldConstraint("foo")
	// var jsonRaw map[string]interface{}
	// json.Unmarshal([]byte(`{ "bar": 1 }`), &jsonRaw)
	//
	// problems := []string{}
	// cut.Check(jsonRaw, "a.b.c", &problems)
	//
	// if len(problems) != 1 {
	// 	t.Errorf("Expected problems size to be 1, got %d", len(problems))
	// }
	//
	// cut = NewHasFieldConstraint("bar")
	// json.Unmarshal([]byte(`{ "bar": 1 }`), &jsonRaw)
	//
	// problems = []string{}
	// cut.Check(jsonRaw, "a.b.c", &problems)
	//
	// if len(problems) != 0 {
	// 	t.Errorf("Expected problems size to be 0, got %d", len(problems))
	// }
}

func TestNonEmptyConstraint(t *testing.T) {
	// cut := NewNonEmptyConstraint("foo")
	// var jsonRaw map[string]interface{}
	// json.Unmarshal([]byte(`{ "bar": 1 }`), &jsonRaw)
	//
	// problems := []string{}
	// cut.Check(jsonRaw, "a.b.c", &problems)
	//
	// if len(problems) != 0 {
	// 	t.Errorf("Expected problems size to be 0, got %d", len(problems))
	// }
	//
	// json.Unmarshal([]byte(`{ "foo": 1 }`), &jsonRaw)
	//
	// problems = []string{}
	// cut.Check(jsonRaw, "a.b.c", &problems)
	//
	// if len(problems) != 0 {
	// 	t.Errorf("Expected problems size to be 0, got %d", len(problems))
	// }
	//
	// json.Unmarshal([]byte(`{ "foo": [ 1 ] }`), &jsonRaw)
	//
	// problems = []string{}
	// cut.Check(jsonRaw, "a.b.c", &problems)
	//
	// if len(problems) != 0 {
	// 	t.Errorf("Expected problems size to be 0, got %d", len(problems))
	// }
	//
	// json.Unmarshal([]byte(`{ "foo": [ ] }`), &jsonRaw)
	//
	// problems = []string{}
	// cut.Check(jsonRaw, "a.b.c", &problems)
	//
	// if len(problems) != 1 {
	// 	t.Errorf("Expected problems size to be 1, got %d", len(problems))
	// }
}
