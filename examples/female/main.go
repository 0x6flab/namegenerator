package main

import (
	"fmt"

	"github.com/0x6flab/namegenerator"
)

func main() {
	generator := namegenerator.NewNameGenerator().WithGender("female")
	name := generator.Generate()
	fmt.Println(name)
}
