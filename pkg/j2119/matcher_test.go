package j2119

import (
	"testing"
	"golang.org/x/exp/slices"
)
func TestShouldFindEachOfLines(t *testing.T){
  cut := Matcher{}
  cut.New("message")
  for _, r := range []string{"Pass State", "Task State", "Choice State", "Parallel State", "Succeed State", "Fail State", "Task State"}{
    cut.AddRole(r)
  }
  
  eachOfLines := []string{
    "Each of a Pass State, a Task State, a Choice State, and a Parallel State MAY have a boolean field named \"End\".",
    "Each of a Succeed State and a Fail State is a \"Terminal State\".",
    "Each of a Task State and a Parallel State MAY have an object-array field named \"Catch\"; each member is a \"Catcher\".",
  }

  for _, eol := range eachOfLines{
    if !cut.eachOfMatch.MatchString(eol){
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
  for _, v := range []string{"Seconds", "SecondsPath", "Timestamp", "TimestampPath"}{
    if !slices.Contains(sl, v){
      t.Fatal("Field list capture group did not contain value", v)
    }
  }
  if len(sl) != 4 {
    t.Fatal("Field list capture group should have 4 values but has", len(sl))
  }
}

