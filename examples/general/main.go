package main

import (
	"fmt"

	"github.com/0x6flab/namegenerator"
)

func main() {
	generator := namegenerator.NewNameGenerator()
	name := generator.Generate()
	fmt.Println(name)
}
