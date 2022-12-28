package matcher

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/catdevman/go-j2119/internal/oxford"
)
 const MUST = "(?P<modal>MUST|MAY|MUST NOT)"
const S = `"[^"]*`
const V = `\S+`
const CHILD_ROLE = `;\s+((its\s+(?P<child_type>value))|` +
                 `each\s+(?P<child_type>field|element))` +
                 `\s+is\s+an?\s+` +
                 `"(?P<child_role>[^"]+)"`

type Matcher struct{
  initialized bool
  eachOfMatch *regexp.Regexp
  roleDefMatch *regexp.Regexp
  constraintStart *regexp.Regexp
  constraintMatch *regexp.Regexp
  onlyOneStart *regexp.Regexp
  onlyOneMatch *regexp.Regexp
  roleMatcher string
  typeRegex string
  strs string
  predicate string
  roles []string
}

func (m *Matcher) New(root string){
  m.makeTypeRegex()
  m.constants()
  m.roles = []string{}
  m.roles = append(m.roles, root)
  m.reconstruct()
}

func (m *Matcher) constants() {
  if !m.initialized {
    m.initialized = true
    opts := oxford.Options{CaptureName: "strings"}
    ox := oxford.Oxford{}
    m.strs = ox.Re(S, opts)
    enum := fmt.Sprintf(`one\s+of\s+%s`, m.strs)
    relation := fmt.Sprintf("((?P<relation>%s)\\s+)", strings.Join([]string{
      "",
      "equal to",
      "greater than",
      "less than",
      "greater than or equal to",
      "less than or equal to",
    }, "|"))
    relational := fmt.Sprintf("%s(?P<target>%s|%s)", relation, S, V)
    m.predicate = fmt.Sprintf("(%s|%s)", relational, enum)
  }
}

func (m *Matcher) reconstruct(){
  m.makeTypeRegex()
  ox := oxford.Oxford{}
  excludedRoles := `not\\s+` + ox.Re(
    m.roleMatcher,
    oxford.Options{
      CaptureName: "excluded",
      UseArticle: true,
    }) + "\\s+"
  conditional := `which\\s+is\\s+` + excludedRoles
  cStart := fmt.Sprintf(`^An?\s+(?P<role>%s)\s+(%s)?%s\s+have\s+an?\s+`, m.roleMatcher, conditional, MUST)
  fieldList := "one\\s+of\\s+" + ox.Re(`"[^"]+"`, oxford.Options{
    CaptureName: "field_list",
  })
  cMatch := fmt.Sprintf(`%s((?P<type>%s)\\s+)?field\\s+named\\s+((\"(?P<field_name>[^\"]+)\")|(%s))(\s+whose\s+value\s+MUST\s+be\s+%s)?(%s)?\.`, cStart, m.typeRegex,fieldList, m.predicate, CHILD_ROLE)
  ooStart := fmt.Sprintf(`^An?\s+(?P<role>%s)\s+%s\s+have\s+only\s+`, m.roleMatcher, MUST)
  ooFieldList := fmt.Sprintf(`one\\s+of\\s+%s`, ox.Re(`"[^"]+"`, oxford.Options{
    CaptureName: "field_list",
    Connector: "and",
  }))

  ooMatch := ooStart + ooFieldList

  valMatch := `whose\\s+\"(?P<fieldtomatch>[^\"]+)\"\\s+field\'s\\s+value\\s+is\\s+(?P<valtomatch>(\"[^\"]*\")|([^\"\\s]\\S+))`
  withAMatch := `with\\s+an?\\s+\"(?P<with_a_field>[^\"]+)\"\\s+field\\s`
  rdMatch := fmt.Sprintf(`^An?\s+(?P<role>%s)\s+((?P<val_match_present>%s)|(%s))?is\\s+an?\\s+\"(?P<newrole>[^\"]*)\"\\.\\s*$`, m.roleMatcher, valMatch, withAMatch)
  
  m.roleDefMatch = regexp.MustCompile(rdMatch)
  
  m.constraintStart = regexp.MustCompile(cStart)
  m.constraintMatch = regexp.MustCompile(cMatch)

  m.onlyOneStart = regexp.MustCompile(ooStart)
  m.onlyOneMatch = regexp.MustCompile(ooMatch)

  eoMatch := fmt.Sprintf(`^Each\\s+of\\s%s\\s+(?P<trailer>.*)$`, ox.Re(m.roleMatcher, oxford.Options{
    CaptureName: "each_of",
    UseArticle: true,
    Connector: "and",
  }))

  m.eachOfMatch = regexp.MustCompile(eoMatch)
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
  m.roles = append(m.roles, role)
  m.roleMatcher = strings.Join(m.roles, "|")
  m.reconstruct()
}

func (m *Matcher) IsOnlyOneMatchLine(line string) bool{
  return m.onlyOneStart.MatchString(line)
}

func (m *Matcher) IsConstraintLine(line string) bool{
  return m.constraintStart.MatchString(line)
}

func (m *Matcher) IsRoleDefLine(line string) bool{
  re := regexp.MustCompile(`is\s+an?\s+"[^"]*"\.\s*$`)
  return re.MatchString(line)
}
