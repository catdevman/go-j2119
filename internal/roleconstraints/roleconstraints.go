package roleconstraints

type RoleConstraints struct{
  Constraints map[string][]string
}

func (r *RoleConstraints) Add(role, constraint string){
  if _, ok := r.Constraints[role]; !ok{
    r.Constraints[role] = []string{}
  }
  r.Constraints[role] = append(r.Constraints[role], constraint)

}

func (r *RoleConstraints) Get(role string) []string {
  if v, ok := r.Constraints[role]; ok{
    return v
  }
  return []string{}
}


