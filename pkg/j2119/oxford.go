// Oxford - all functionality completed
package j2119

import (
	"fmt"
	"regexp"
	"strings"
)

type OxfordOptions struct {
	UseArticle  bool
	CaptureName string
	Connector   string
}

type Oxford struct{}

const BASIC = "(?P<CAPTURE>X((((,\\s+X)+,)?)?\\s+or\\s+X)?)"

func (o *Oxford) Re(particle string, opts OxfordOptions) string {
	basic := strings.Split(BASIC, "X")
	hasCapture := basic[0]
	inter := basic[1]
	hasConnector := basic[2]
	last := basic[3]
	if opts.Connector != "" {
		hasConnector = strings.ReplaceAll(hasConnector, "or", opts.Connector)
	}
	if opts.UseArticle {
		particle = fmt.Sprintf("an?\\s+(%s)", particle)
	} else {
		particle = fmt.Sprintf("(%s)", particle)
	}

	if opts.CaptureName != "" {
		hasCapture = strings.ReplaceAll(hasCapture, "CAPTURE", opts.CaptureName)
	} else {
		hasCapture = strings.ReplaceAll(hasCapture, "?P<CAPTURE>", "")
	}

	return strings.Join([]string{hasCapture, inter, hasConnector, last}, particle)
}

func (o *Oxford) BreakStringList(list string) []string {
	pieces := make([]string, 0)
	re := regexp.MustCompile(`^[^"]*"([^"]*)"`)
	for {
		match := re.FindStringSubmatch(list)
		if match == nil {
			break
		}
		pieces = append(pieces, match[1])
		list = list[len(match[0]):]
	}
	return pieces
}

func (o *Oxford) BreakRoleList(matcher Matcher, list string) []string {
	pieces := []string{}
	re := regexp.MustCompile(fmt.Sprintf(`^an?\s+(%s)(,\s+)?`, matcher.RoleMatcher))
	for re.MatchString(list) {
		submatches := re.FindStringSubmatch(list)
		pieces = append(pieces, submatches[1])
		list = strings.TrimPrefix(list, submatches[0])
	}
	if re := regexp.MustCompile(fmt.Sprintf(`^\s*(and|or)\s+an?\s+(%s)`, matcher.RoleMatcher)); re.MatchString(list) {
		submatches := re.FindStringSubmatch(list)
		pieces = append(pieces, submatches[2])
	}
	return pieces
}
