package j2119

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Validator struct {
	Parsed string
	Root   string
	parser Parser
}

func (v *Validator) Init(reader io.Reader) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
	}
	fmt.Println(data)
}

func (v *Validator) Validate(reader io.Reader) []string {
	return []string{}
}
