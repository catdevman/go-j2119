package j2119

import (
	"os"
)

type Validator struct {
	Parsed string
	Root   string
	parser Parser
}

func (v *Validator) Init(schema *os.File) {
	defer schema.Close()
	v.parser.New(schema)
}

func (v *Validator) Validate(source *os.File) []string {
	return []string{}
}
