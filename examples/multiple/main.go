package main

import (
	"fmt"

	"github.com/0x6flab/namegenerator"
)

func main() {
	generator := namegenerator.NewNameGenerator()
	name := generator.GenerateNames(10)
	fmt.Println(name)
}
