// Copyright (c) 2023 0x6flab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
// http://www.apache.org/licenses/LICENSE-2.0

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

	details := generator.GenerateNameAndEmail()
	fmt.Printf("Names and Emails: Name: %s, Email: %s \n", details.Name, details.Email)

	generator = namegenerator.NewGenerator().WithRandomString(5)
	name = generator.Generate()
	fmt.Println(name)

	names = generator.GenerateMultiple(10)
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

	generator = namegenerator.NewGenerator()
	name = generator.Generate(namegenerator.WithGender(namegenerator.Male), namegenerator.WithPrefix("Dr. "), namegenerator.WithSuffix(" Jr."), namegenerator.WithRandomString(5))
	fmt.Println(name)

	names = generator.GenerateMultiple(10, namegenerator.WithGender(namegenerator.Male), namegenerator.WithPrefix("Dr. "), namegenerator.WithSuffix(" Jr."), namegenerator.WithRandomString(5))
	fmt.Println(names)
}
