package j2119

import (
	"bufio"
	"os"
  "fmt"
	"github.com/catdevman/go-j2119/internal"
)

type Validator struct{
  parser internal.Parser
}

func (v *Validator) Init(schema *os.File){
  defer schema.Close()
  scanner := bufio.NewScanner(schema)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
}

func (v *Validator) Validate(source *os.File) []string{
  return []string{}
}
