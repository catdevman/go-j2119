package j2119

import (
  "fmt"
  "os"
)

type Validator struct{}

func (v *Validator) New(schema os.File){
  fmt.Println(schema)
}
