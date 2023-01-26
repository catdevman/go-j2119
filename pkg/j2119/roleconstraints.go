package j2119 

type RoleConstraints struct{
  Constraints map[string][]Constraint
}

func (r *RoleConstraints) New(){}

func (r *RoleConstraints) Add(role string, constraint Constraint){ 
  if _, ok := r.Constraints[role]; !ok{
    r.Constraints[role] = []Constraint{}
  }
  r.Constraints[role] = append(r.Constraints[role], constraint)

}

func (r *RoleConstraints) Get(role string) []Constraint {
  if v, ok := r.Constraints[role]; ok{
    return v
  }
  return []Constraint{}
}


