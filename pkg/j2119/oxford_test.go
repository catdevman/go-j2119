package j2119 

//import (
//	"regexp"
//	"testing"
//)
//
//func TestItShouldShowTheUnderlayingPatternWorking(t *testing.T) {
//  re := regexp.MustCompile(`^` + BASIC + `$`)
//
//  if !re.MatchString("X"){
//    t.Fail()
//  }
//
//  if !re.MatchString("X or X") {
//    t.Fail()
//  }
//
//  if !re.MatchString("X, X, or X"){
//    t.Fail()
//  }
//
//  if !re.MatchString("X, X, X, or X") {
//    t.Fail()
//  }
//}
//
//func TestItShouldDoANoArticleNoCaptureNoConnectMatch(t *testing.T) {
//  targets := []string{
//    "a",
//    "a or aa",
//    "a, aa, or aaa",
//    "a, aa, aaa, or aaaa",
//  }
//  oxford := Oxford{}
//  opts := Options{}
//  re := regexp.MustCompile("^" + oxford.Re("a+", opts) + "$")
//
//  for _, target := range targets {
//    if !re.MatchString(target){
//      t.Fail()
//    }
//  }
//}
//
//func TestItShouldDoOneWithCaptureArticlesConnector(t *testing.T){
//  targets := []string{
//    `an "asdg"`,
//    `a "foij2pe" and an "aiepw"`,
//    `an "alkvm 2", an "ap89wf", and a " lfdj a fddalfkj"`,
//    `an "aj89peww", a "", an "aslk9 ", and an "x"`,
//  }
//  oxford := Oxford{}
//  opts := Options{
//    Connector: "and",
//    UseArticle: true,
//    CaptureName: "capture_me",
//  }
//  re := oxford.Re(`"([^"]*)"`, opts)
//  cut := regexp.MustCompile("^" + re + "$")
//
//  for _, target := range targets{
//    if !cut.MatchString(target) {
//      t.Fail()
//    }
//  }
//}
//
//func TestItShouldProperlyBreakUpARoleList(t *testing.T) {
//  OXFORD_LIST := []string{
//    "an R2",
//    "an R2 or an R3",
//    "an R2, an R3, or an R4",
//  }
//  WANTED_PIECES := [][]string{
//    {
//      "R2",
//    },
//    {
//      "R2",
//      "R3",
//    },
//    {
//      "R2",
//      "R3",
//      "R4",
//    },
//  }
//  oxford := Oxford{}
//  for i, list := range OXFORD_LIST{
//    if !sameStringSlices(oxford.BreakRoleList("R1", list), WANTED_PIECES[i]){
//      t.Fail()
//    }
//  }
//}
//
//func sameStringSlices(a []string, b []string) bool{
//  if len(a) != len(b){
//    return false
//  }
//
//  for i, v := range a{
//    if v != b[i]{
//      return false
//    }
//  }
//  return true
//}
