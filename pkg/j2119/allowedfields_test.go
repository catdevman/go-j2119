// All tests completed
package j2119

import (
	"fmt"
	"testing"
)

func TestReturnsTrueWhenAppropriate(t *testing.T) {
	cut := NewAllowedFields()
	cut.SetAllowed("foo", "bar")
	fmt.Println("test")
	if !cut.Allowed([]string{"foo"}, "bar") {
		t.Fail()
	}

	if !cut.Allowed([]string{"bar", "baz", "foo"}, "bar") {
		t.Fail()
	}
}

func TestReturnsFalseWhenAppropriate(t *testing.T) {
	cut := NewAllowedFields()
	cut.SetAllowed("foo", "bar")

	if cut.Allowed([]string{"foo"}, "baz") {
		t.Fail()
	}

	if cut.Allowed([]string{"bar", "baz", "foo"}, "baz") {
		t.Fail()
	}
}

func TestWorksWithStrangeQueries(t *testing.T) {
	cut := NewAllowedFields()
	cut.SetAllowed("foo", "bar")

	if cut.Allowed([]string{"boo"}, "baz") {
		t.Fail()
	}

	if cut.Allowed([]string{}, "baz") {
		t.Fail()
	}
}
