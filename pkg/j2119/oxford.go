package j2119 

import (
	"fmt"
	"regexp"
	"strings"
)

type Options struct{
  UseArticle bool
  CaptureName string
  Connector string
}

type Oxford struct{}

const BASIC  = "(?P<CAPTURE>X((((,\\s+X)+,)?)?\\s+or\\s+X)?)"

func (o *Oxford) Re(particle string, opts Options) string{
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

func (o *Oxford) BreakStringList(list string) []string{
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

func (o *Oxford) BreakRoleList(matcher string, list string) []string {
    pieces := make([]string, 0)
    re := regexp.MustCompile(`^an?\s+` + matcher + `(,\s+)?`)
    fmt.Println(re.FindAllString(list, -1))
//    for {
//        match := re.FindString(list)
//        fmt.Println(match)
//        if match == nil {
//            break
//        }
//        pieces = append(pieces, match[1])
//        list = list[len(match[0]):]
//    }
    re2 := regexp.MustCompile(`^\s*(and|or)\s+an?\s+`+matcher)
    if matches := re2.FindAllString(list, -1); len(matches) >= 3  {
        pieces = append(pieces, matches[2])
    }
  
    return pieces
}
