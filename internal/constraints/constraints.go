package constraints

import (
  "fmt"
  "strings"
  "github.com/catdevman/go-j2119/internal/conditional"
)
type Constraint struct{
  Conditions []conditional.RoleNotPresentCondition
}

type OnlyOneOfConstraint struct{
  Constraint
  Fields []string
}

func (o *OnlyOneOfConstraint) New(fields []string){
  o.Fields = append(o.Fields, fields...)
}

func (o *OnlyOneOfConstraint) Check(node map[string]interface{}, path string, problems []string){
  var count int
  for k := range node{
    for _, f := range o.Fields{
      if k == f {
        count = count + 1
      }
    }
    if count > 1 {
      problems = append(problems, fmt.Sprintf("%s may have only one of %s", path, strings.Join(o.Fields, ",")))
    }
  } 
}

type NonEmptyConstraint struct{
  Constraint
  name string
}

func (n *NonEmptyConstraint) New(name string){
  n.name = name
}

func (n *NonEmptyConstraint) String() string{
  conds := ""
  if len(n.Constraint.Conditions) > 0{
    conds = fmt.Sprintf(" %s conditions", len(n.Constraint.Conditions))
  }

  return fmt.Sprintf("<Array field %s should not be empty%s>", n.name, conds)
}

func (n *NonEmptyConstraint) Check(node map[string]interface{}, path string, problems []string){
  n, ok := node[n.name]
  if ok {
    v, ok := n.([]interface{})
    if ok && len(v) == 0 {
      problems = append(problems, fmt.Sprintf("%s.%s is empty, non-empty required", path, n.name))
    }
  }
}



