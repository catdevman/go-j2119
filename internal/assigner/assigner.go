package assigner

import(
  "github.com/catdevman/go-j2119/internal/roleconstraints"
  "github.com/catdevman/go-j2119/internal/rolefinder"
  "github.com/catdevman/go-j2119/internal/matcher"
  "github.com/catdevman/go-j2119/internal/allowedfields"
)

type Assigner struct{}

func (a *Assigner) New(roleconstraints.RoleConstraints, rolefinder.RoleFinder, matcher.Matcher, allowedfields.AllowedFields){

}
