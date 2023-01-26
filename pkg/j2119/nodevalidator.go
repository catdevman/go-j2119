package j2119 

type NodeValidator struct{
  parser Parser
}

func (nv *NodeValidator) Initialize(p Parser){
  nv.parser = p
}

func (nv *NodeValidator) ValidateNode(node map[string]interface{}, path string, roles []string, problems []string){
//  nv.FindMoreRoles(node, roles)

 // for _, role := range roles{
 //   constraints := nv.GetConstraints(role)
 //   for _, constraint := range constraints{

 //   }
 // }
}
