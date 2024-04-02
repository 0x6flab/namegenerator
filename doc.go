// Copyright (c) 2023 0x6flab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
// http://www.apache.org/licenses/LICENSE-2.0

// Package namegenerator provides a simple way to generate random names.
//
// Example to generate general names:
//
//	package main
//
//	import (
//
//		"fmt"
//
//		"github.com/0x6flab/namegenerator"
//
//	)
//
//	func main() {
//		generator := namegenerator.NewGenerator()
//		name := generator.Generate()
//		fmt.Println(name)
//
//		names := generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewGenerator().WithGender(namegenerator.Male)
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewGenerator().WithGender(namegenerator.Female)
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewGenerator().WithGender(namegenerator.NonBinary)
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewGenerator().WithPrefix("Dr. ")
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewGenerator().WithSuffix(" Jr.")
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewGenerator().WithPrefix("Dr. ").WithSuffix(" Jr.")
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewGenerator().WithGender(namegenerator.Male).WithPrefix("Dr. ").WithSuffix(" Jr.")
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//	}
package namegenerator
