package matcher

import(
  "fmt"
  "regexp"
  "strings"
)

type Matcher struct{
  roleMatcher string
  eachOfMatch *regexp.Regexp
  typeRegex string
  roles []string
}

func (m *Matcher) New(root string){
  m.makeTypeRegex()
  //TODO: build constants
  //TODO: set roles to []string
  //TODO: add role
  //TODO: reconstruct (rebuild regex)
}

func (m *Matcher) reconstruct(){
  //TODO: build types regex
}

func (m *Matcher) makeTypeRegex(){
  types := []string{
    "array",
    "object",
    "string",
    "boolean",
    "numeric",
    "integer",
    "float",
    "timestamp",
    "JSONPath",
    "referencePath",
    "URI",
  }
  numberTypes := []string{
    "float",
    "integer",
    "numeric",
  }
  numberMods := []string{
    "positive",
    "negative",
    "nonnegative",
  }
  for _, ty := range numberTypes{
    for _, mod := range numberMods{
      types = append(types, fmt.Sprintf("%s-%s", mod, ty))
    }
  }

  arrayTypes := []string{}
  nonemptyArrayTypes := []string{}
  nullableTypes := []string{}
  for _, t := range types{
    arrayTypes = append(arrayTypes, fmt.Sprintf("%s-array", t))
    nonemptyArrayTypes = append(nonemptyArrayTypes, fmt.Sprintf("nonempty-%s", t))
    nullableTypes = append(nullableTypes, fmt.Sprintf("nullable-%s", t))
  }
  types = append(types, arrayTypes...)
  types = append(types, nonemptyArrayTypes...)
  types = append(types, nullableTypes...)

  m.typeRegex = strings.Join(types, "|")  
}

func (m *Matcher) AddRole(role string){
  //TODO: add role to roles slice string
  //TODO: set roleMatcher to roles join as "|"
  //TODO: reconstruct (rebuild regex)
}
