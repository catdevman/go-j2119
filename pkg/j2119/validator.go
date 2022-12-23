package j2119

import (
  "os"
	"github.com/catdevman/go-j2119/internal/parser"
)

type Validator struct{
  Parsed string
  Root string
  parser parser.Parser
}

func (v *Validator) Init(schema *os.File){
  defer schema.Close()
  v.parser.New(schema)
}

func (v *Validator) Validate(source *os.File) []string{
  return []string{}
}
