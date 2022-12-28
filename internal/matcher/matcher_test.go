package matcher

import (
  "testing"
)
func TestShouldHandleOnlyOneOfLines(t *testing.T) {
  line := `A x MUST have only one of "Seconds", "SecondsPath", "Timestamp", and "TimestampPath"`
  cut := Matcher{}
  cut.New("x")
  if !cut.IsOnlyOneMatchLine(line) {
    t.Fail()
  }
  m := cut.onlyOneMatch.FindString(line)

}
