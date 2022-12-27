package nodevalidator

import "github.com/catdevman/go-j2119/internal/parser"



type NodeValidator struct{
  parser parser.Parser
}

func (nv *NodeValidator) Initialize(p parser.Parser){
  nv.parser = p
}

func (nv *NodeValidator) ValidateNode(node map[string]interface{}, path string, roles []string, problems []string){
  nv.parser.FindMoreRoles(node, roles)

 // for _, role := range roles{
 //   constraints := nv.parser.GetConstraints(role)
 //   for _, constraint := range constraints{

 //   }
 // }
}
