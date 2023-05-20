package j2119

import (
    "fmt"
    "strconv"
    "regexp"
)


func Deduce(val string) (interface{}, error) {
    switch {
    case regexp.MustCompile(`^"(.*)"$`).MatchString(val):
        return val[1 : len(val)-1], nil
    case val == "true":
        return true, nil
    case val == "false":
        return false, nil
    case val == "null":
        return nil, nil
    case regexp.MustCompile(`^\d+$`).MatchString(val):
        i, err := strconv.Atoi(val)
        if err != nil {
            return nil, fmt.Errorf("error converting string to int: %w", err)
        }
        return i, nil
    default:
        f, err := strconv.ParseFloat(val, 64)
        if err != nil {
            return nil, fmt.Errorf("error converting string to float: %w", err)
        }
        return f, nil
    }
}
