package main

import (
	"os"
	"testing"

	"github.com/catdevman/go-j2119/pkg/j2119"
)

func TestMain(t *testing.T){
    validator := j2119.Validator{
    }
    file, _ := os.Open("test.json")
    defer file.Close()
    validator.Init(file)
    validator.Validate(file)
}
