package j2119

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestShouldFindEachOfLines(t *testing.T) {
	cut := Matcher{}
	cut.New("message")
	for _, r := range []string{"Pass State", "Task State", "Choice State", "Parallel State", "Succeed State", "Fail State", "Task State"} {
		cut.AddRole(r)
	}

	eachOfLines := []string{
		"Each of a Pass State, a Task State, a Choice State, and a Parallel State MAY have a boolean field named \"End\".",
		"Each of a Succeed State and a Fail State is a \"Terminal State\".",
		"Each of a Task State and a Parallel State MAY have an object-array field named \"Catch\"; each member is a \"Catcher\".",
	}

	for _, eol := range eachOfLines {
		if !cut.eachOfMatch.MatchString(eol) {
			t.Fatal("Should match each of line", eol)
		}
	}
}

func TestShouldHandleOnlyOneOfLines(t *testing.T) {
	line := `A x MUST have only one of "Seconds", "SecondsPath", "Timestamp", and "TimestampPath"`
	cut := Matcher{}
	cut.New("x")
	if !cut.IsOnlyOneMatchLine(line) {
		t.Fail()
	}
	results := reSubMatchMap(cut.onlyOneMatch, line)
	if v, ok := results["role"]; ok {
		if v != "x" {
			t.Fatal("The role should have been x but was", v)
		}
	} else {
		t.Fatal("Failed to find named capture group for role")
	}

	s, ok := results["field_list"]
	if !ok {
		t.Fatal("Failed to fild named capture group field_list")
	}
	ox := Oxford{}
	sl := ox.BreakStringList(s)
	for _, v := range []string{"Seconds", "SecondsPath", "Timestamp", "TimestampPath"} {
		if !slices.Contains(sl, v) {
			t.Fatal("Field list capture group did not contain value", v)
		}
	}
	if len(sl) != 4 {
		t.Fatal("Field list capture group should have 4 values but has", len(sl))
	}
}

func TestShouldDisassembleEachofLines(t *testing.T) {
	splitEachofLines := [][]string{
		{
			"A Pass State MAY have a boolean field named \"End\".",
			"A Task State MAY have a boolean field named \"End\".",
			"A Choice State MAY have a boolean field named \"End\".",
			"A Parallel State MAY have a boolean field named \"End\".",
		},
		{
			"A Succeed State is a \"Terminal State\".",
			"A Fail State is a \"Terminal State\".",
		},
		{
			"A Task State MAY have an object-array field named \"Catch\"; each member is a \"Catcher\".",
			"A Parallel State MAY have an object-array field named \"Catch\"; each member is a \"Catcher\".",
		},
	}

	eachOfLines := []string{
		"Each of a Pass State, a Task State, a Choice State, and a Parallel State MAY have a boolean field named \"End\".",
		"Each of a Succeed State and a Fail State is a \"Terminal State\".",
		"Each of a Task State and a Parallel State MAY have an object-array field named \"Catch\"; each member is a \"Catcher\".",
	}

	oxford := Oxford{}

	cut := Matcher{}
	cut.New("message")

	for _, r := range []string{"Pass State", "Task State", "Choice State", "Parallel State", "Succeed State", "Fail State", "Task State"} {
		cut.AddRole(r)
	}

	for idx, line := range eachOfLines {
		wanted := splitEachofLines[idx]
		for _, oneLine := range oxford.BreakRoleList(cut.roleMatcher, line) {
			if !slices.Contains(wanted, oneLine) {
				t.Fail()
			}
		}
	}
}

func TestSpotRoleDefLines(t *testing.T) {
	rdlines := []string{
		"A State whose \"End\" field's value is true is a \"Terminal State\".",
		"Each of a Secceed State and a Fail State is a \"Terminal State\".",
		"A Choice Rule with a \"Variable\" field is a \"Comparison\".",
	}

	cut := Matcher{}
	cut.New("message")

	for _, rdline := range rdlines {
		if !cut.IsRoleDefLine(rdline) {
			t.Fail()
		}
	}
}

func TestMatchValueBasedRoleDefs(t *testing.T) {
	valueBasedRoleDefs := []string{
		"A State whose \"End\" field's value is true is a \"Terminal State\".",
		"A State whose \"Comment\" field's value is \"Hi\" is a \"Frobble\".",
		"A State with a \"Foo\" field is a \"Bar\".",
	}

	cut := Matcher{}
	cut.New("State")

	for _, v := range valueBasedRoleDefs {
		if !cut.roleDefMatch.MatchString(v) {
			t.Fail()
		}
	}

	m1 := reSubMatchMap(cut.roleDefMatch, valueBasedRoleDefs[0])
	if v, ok := m1["role"]; !ok || v != "State" {
		t.Fatal("role not set correctly")
	}

	if v, ok := m1["fieldtomatch"]; !ok || v != "End" {
		t.Fatal("fieldtomatch not set correctly")
	}

	if v, ok := m1["valtomatch"]; !ok || v != "true" {
		t.Fatal("valtomatch not set correctly")
	}

	if v, ok := m1["newrole"]; !ok || v != "Terminal State" {
		t.Fatal("newrole not set correctly")
	}

	if v, ok := m1["val_match_present"]; !ok || v == "" {
		t.Fatal("val_match_present is not set correctly")
	}

	m2 := reSubMatchMap(cut.roleDefMatch, valueBasedRoleDefs[1])
	if v, ok := m2["role"]; !ok || v != "State" {
		t.Fatal("role not set correctly")
	}

	if v, ok := m2["fieldtomatch"]; !ok || v != "Comment" {
		t.Fatal("fieldtomatch is not set correctly")
	}

	if v, ok := m2["valtomatch"]; !ok || v != `"Hi"` {
		t.Fatal("valtomatch is not set correctly")
	}

	if v, ok := m2["newrole"]; !ok || v != "Frobble" {
		t.Fatal("newrole is not set correctly")
	}

	if v, ok := m2["val_match_present"]; !ok || v == "" {
		t.Fatal("val_match_present is not set correctly")
	}

	m3 := reSubMatchMap(cut.roleDefMatch, valueBasedRoleDefs[2])

	if v, ok := m3["role"]; !ok || v != "State" {
		t.Fatal("role not set correctly")
	}

	if v, ok := m3["newrole"]; !ok || v != "Bar" {
		t.Fatal("newrole not set correctly")
	}

	if v, ok := m3["with_a_field"]; !ok || v != "Foo" {
		t.Fatal("with_a_field is not set correctly")
	}
}

func TestMatchIsARoleDefs(t *testing.T) {
	cut := Matcher{}
	cut.New("Foo")

	if !cut.roleDefMatch.MatchString(`A Foo is a "Bar".`) {
		t.Fatal("roleDefMatch did not correctly match")
	}
}

func TestParseIsARoleDefs(t *testing.T) {
	cut := Matcher{}
	cut.New("Foo")
	cut.AddRole("Bar")

	m1 := reSubMatchMap(cut.roleDefMatch, `A Foo is a "Bar".`)
	if v, ok := m1["val_match_present"]; !ok || v != "" {
		t.Fatal("val_match_present should not be set")
	}
}

func TestParseValueBasedRoleDefs(t *testing.T) {
	cut := Matcher{}
	cut.New("State")

	valueBasedRoleDefs := []string{
		"A State whose \"End\" field's value is true is a \"Terminal State\".",
		"A State whose \"Comment\" field's value is \"Hi\" is a \"Frobble\".",
		"A State with a \"Foo\" field is a \"Bar\".",
	}
	m1 := reSubMatchMap(cut.roleDefMatch, valueBasedRoleDefs[0])
	if v, ok := m1["role"]; !ok || v != "State" {
		t.Fatal("role not set correctly")
	}

	if v, ok := m1["fieldtomatch"]; !ok || v != "End" {
		t.Fatal("fieldtomatch is not set correctly")
	}

	if v, ok := m1["valtomatch"]; !ok || v != "true" {
		t.Fatal("valtomatch is not set correctly")
	}

	if v, ok := m1["newrole"]; !ok || v != "Terminal State" {
		t.Fatal("newrole is not set correctly")
	}

	m2 := reSubMatchMap(cut.roleDefMatch, valueBasedRoleDefs[1])

	if v, ok := m2["role"]; !ok || v != "State" {
		t.Fatal("role is not set correctly")
	}

	if v, ok := m2["fieldtomatch"]; !ok || v != "Comment" {
		t.Fatal("fieldtomatch is not set correctly")
	}

	if v, ok := m2["valtomatch"]; !ok || v != `"Hi"` {
		t.Fatal("valtomatch is not set correctly")
	}

	if v, ok := m2["newrole"]; !ok || v != "Frobble" {
		t.Fatal("newrole is not set correctly")
	}
}

var lines []string = []string{
	`A message MUST have an object field named "States"; each field is a "State".`,
	`A message MUST have a negative-integer-array field named "StartAt".`,
	`A message MAY have a string-array field named "StartAt"`,
	`A message MUST NOT have a field named "StartAt"`,
	`A message MUST have a field named one of "StringEquals", "StringLessThan", "StringGreaterThan", "StringLessThanEquals", "StringGreaterThanEquals", "NumericEquals", "NumericLessThan", "NumericGreaterThan", "NumericLessThanEquals", "NumericGreaterThanEquals", "BooleanEquals", "TimestampEquals", "TimestampLessThan", "TimestampGreaterThan", "TimestampLessThanEquals", or "TimestampGreaterThanEquals"`,
}

func TestSpotASimpleConstraintLine(t *testing.T) {
	cut := Matcher{}
	cut.New("message")
	for _, line := range lines {
		if !cut.IsConstraintLine(line) {
			t.Fatal("Did not recognize constraint line")
		}
	}
}

func TestSpotASimpleConstraintLineWithNewRoles(t *testing.T) {
	cut := Matcher{}
	cut.New("message")
	lines2 := []string{}
	for _, line := range lines {
		lines2 = append(lines2, strings.ReplaceAll(line, "message", "avatar"))
	}
	cut.AddRole("avatar")
	for _, line := range lines2 {
		if !cut.IsConstraintLine(line) {
			t.Fatal("Did not recognize constraint line")
		}
	}
}

var condLines = []string{
	`An R1 MUST have an object field named "States"; each field is a "State".`,
	`An R1 which is not an R2 MUST have an object field named "States"; each field is a "State".`,
	`An R1 which is not an R2 or an R3 MUST NOT have a field named "StartAt".`,
	`An R1 which is not an R2, an R3, or an R4 MUST NOT have a field named "StartAt".`,
}

func TestCatchAConditionalOnAConstraint(t *testing.T) {
	excludes := []string{
		"",
		"an R2",
		"an R2 or an R3",
		"an R2, an R3, or an R4",
	}
	cut := Matcher{}
	cut.New("R1")
	cut.AddRole("R2")
	cut.AddRole("R3")
	cut.AddRole("R4")

	for idx, condLine := range condLines {
		if !cut.constraintMatch.MatchString(condLine) {
			t.Fatal("Did not match on constraint")
		}
		m := reSubMatchMap(cut.constraintMatch, condLine)
		if v, ok := m["excluded"]; !ok || v != excludes[idx] {
			t.Log(v, excludes[idx])
			t.Fatal("constraint did not match excluded correctly")
		}
	}
}

func TestMatchAReasonablyComplexConstraint(t *testing.T) {
	cut := Matcher{}
	cut.New("State")
	s := `A State MUST have a string field named "Type" whose value MUST be one of "Pass", "Succeed", "Fail", "Task", "Choice", "Wait", or "Parallel".`
	if !cut.constraintMatch.MatchString(s) {
		t.Fatal("failed to match string on complex constraint")
	}

	cut.AddRole("Retrier")

	s = `A Retrier MAY have a nonnegative-integer field named "MaxAttempts" whose value MUST be less than 99999999.`
	if !cut.constraintMatch.MatchString(s) {
		t.Fatal("failed to match string on complex constraint")
	}
}

func TestMatchAnEnumConstraintObject(t *testing.T) {
	cut := Matcher{}
	cut.New("State")
	s := `A State MUST have a string field named "Type" whose value MUST be one of "Pass", "Succeed", "Fail", "Task", "Choice", "Wait", or "Parallel".`
	con := reSubMatchMap(cut.constraintMatch, s)
	if v, ok := con["role"]; !ok {
		t.Log(v)
		t.Fatal("role is not set correctly")
	}

	if v, ok := con["modal"]; !ok || v != "MUST" {
		t.Fatal("modal is not set correctly")
	}

	if v, ok := con["type"]; !ok || v != "string" {
		t.Fatal("type is not set correctly")
	}

	if v, ok := con["field_name"]; !ok || v != "Type" {
		t.Fatal("field_name is not set correctly")
	}

	if v, ok := con["relation"]; !ok || v != "" {
		t.Fatal("relation is not set correctly")
	}
}
