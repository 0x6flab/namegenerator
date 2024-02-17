// Copyright (c) 2023 0x6flab
//
// SPDX-License-Identifier: GNU GENERAL PUBLIC LICENSE

package main

import (
	"fmt"

	"github.com/0x6flab/namegenerator"
)

func main() {
	generator := namegenerator.NewGenerator()
	name := generator.Generate()
	fmt.Println(name)

	names := generator.GenerateMultiple(10)
	fmt.Println(names)

	generator = namegenerator.NewGenerator().WithGender(namegenerator.Male)
	name = generator.Generate()
	fmt.Println(name)

	names = generator.GenerateMultiple(10)
	fmt.Println(names)

	generator = namegenerator.NewGenerator().WithGender(namegenerator.Female)
	name = generator.Generate()
	fmt.Println(name)

	names = generator.GenerateMultiple(10)
	fmt.Println(names)

	generator = namegenerator.NewGenerator().WithGender(namegenerator.NonBinary)
	name = generator.Generate()
	fmt.Println(name)

	names = generator.GenerateMultiple(10)
	fmt.Println(names)

	generator = namegenerator.NewGenerator().WithPrefix("Dr. ")
	name = generator.Generate()
	fmt.Println(name)

	names = generator.GenerateMultiple(10)
	fmt.Println(names)

	generator = namegenerator.NewGenerator().WithSuffix(" Jr.")
	name = generator.Generate()
	fmt.Println(name)

	names = generator.GenerateMultiple(10)
	fmt.Println(names)

	generator = namegenerator.NewGenerator().WithPrefix("Dr. ").WithSuffix(" Jr.")
	name = generator.Generate()
	fmt.Println(name)

	names = generator.GenerateMultiple(10)
	fmt.Println(names)

	generator = namegenerator.NewGenerator().WithGender(namegenerator.Male).WithPrefix("Dr. ").WithSuffix(" Jr.")
	name = generator.Generate()
	fmt.Println(name)

	names = generator.GenerateMultiple(10)
	fmt.Println(names)
}
