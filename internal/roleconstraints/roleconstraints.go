package roleconstraints

import(
  "github.com/catdevman/go-j2119/internal/constraints"
)

type RoleConstraints struct{
  Constraints map[string][]constraints.Constraint
}

func (r *RoleConstraints) New(){}

func (r *RoleConstraints) Add(role string, constraint constraints.Constraint){ 
  if _, ok := r.Constraints[role]; !ok{
    r.Constraints[role] = []constraints.Constraint{}
  }
  r.Constraints[role] = append(r.Constraints[role], constraint)

}

func (r *RoleConstraints) Get(role string) []constraints.Constraint {
  if v, ok := r.Constraints[role]; ok{
    return v
  }
  return []constraints.Constraint{}
}


