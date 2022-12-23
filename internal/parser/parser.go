package parser

import (
  "fmt"
  "os"
  "bufio"
  "regexp"
  "github.com/catdevman/go-j2119/internal/assigner"
  "github.com/catdevman/go-j2119/internal/matcher"
  "github.com/catdevman/go-j2119/internal/roleconstraints"
  "github.com/catdevman/go-j2119/internal/rolefinder"
  "github.com/catdevman/go-j2119/internal/allowedfields"
)


type Parser struct{
  root string
  haveRoot bool
  failed bool
  assigner assigner.Assigner
  matcher matcher.Matcher
  constraints roleconstraints.RoleConstraints
  finder rolefinder.RoleFinder
  allowedFields allowedfields.AllowedFields
}

func (p *Parser) New(f *os.File){
  ROOT, _ := regexp.Compile(`This\s+document\s+specifies\s+a\s+JSON\s+object\s+called\s+an?\s+"([^"]+)"\.`)
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    line := scanner.Text()
    if ROOT.MatchString(line){
      if p.haveRoot{
        panic("Only one root declaration is allowed") //TODO: return errors??
      } else {
        p.root = ROOT.FindString(line)
        p.matcher.New(p.root)
        p.assigner.New()//p.constraints, p.finder, p.matcher, p.allowedFields)
        p.haveRoot = true
      }
    } else {
      if !p.haveRoot {
        panic("Root declaration must be first") //TODO: return errors??
      } else {
        p.processLine(line)
      }
    }
  }
  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  } 
}

func (p *Parser) processLine(line string){
  fmt.Println(line)
}

