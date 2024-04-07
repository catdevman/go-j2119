package j2119

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Parser struct {
	root          string
	haveRoot      bool
	failed        bool
	assigner      Assigner
	matcher       Matcher
	constraints   RoleConstraints
	finder        RoleFinder
	allowedFields AllowedFields
}

func (p *Parser) New(f *os.File) {
	ROOT, _ := regexp.Compile(`This\s+document\s+specifies\s+a\s+JSON\s+object\s+called\s+an?\s+"([^"]+)"\.`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if ROOT.MatchString(line) {
			if p.haveRoot {
				panic("Only one root declaration is allowed") //TODO: return errors??
			} else {
				p.root = ROOT.FindString(line)
				p.matcher = NewMatcher(p.root)
				p.constraints = NewRoleConstraints()
				p.assigner = NewAssigner(p.constraints, p.finder, p.matcher, p.allowedFields)
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

func (p *Parser) processLine(line string) {
	// p.matcher.IsConstraintLine(p.BuildConstraint(line))
	// fmt.Println(line)
}
