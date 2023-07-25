// deduce tests are complete
package j2119

import(
    "testing"
)

func TestDeduceVal(t *testing.T) {

    if val, err := Deduce(`"x"`); err != nil || val !="x" {
        t.Fatal("was not able to Deduce correctly")
    }

    if val, err := Deduce(`true`); err != nil || val !=true {
        t.Fatal("was not able to Deduce correctly")
    }

    if val, err := Deduce(`false`); err != nil || val != false {
        t.Fatal("was not able to Deduce correctly")
    }

    if val, err := Deduce(`null`); err != nil || val != nil {
        t.Fatal("was not able to Deduce correctly")
    }

    if val, err := Deduce(`234`); err != nil || val != 234 {
        t.Fatal("was not able to Deduce correctly")
    }

    if val, err := Deduce(`1.234`); err != nil || val != 1.234{
        t.Fatal("was not able to Deduce correctly")
    }
}
