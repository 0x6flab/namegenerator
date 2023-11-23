// Copyright (c) 2023 0x6flab
//
// SPDX-License-Identifier: GNU GENERAL PUBLIC LICENSE

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
//		generator := namegenerator.NewNameGenerator()
//		name := generator.Generate()
//		fmt.Println(name)
//
//		names := generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewNameGenerator().WithGender(namegenerator.Male)
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewNameGenerator().WithGender(namegenerator.Female)
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewNameGenerator().WithGender(namegenerator.NonBinary)
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewNameGenerator().WithPrefix("Dr. ")
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewNameGenerator().WithSuffix(" Jr.")
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewNameGenerator().WithPrefix("Dr. ").WithSuffix(" Jr.")
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//
//		generator = namegenerator.NewNameGenerator().WithGender(namegenerator.Male).WithPrefix("Dr. ").WithSuffix(" Jr.")
//		name = generator.Generate()
//		fmt.Println(name)
//
//		names = generator.GenerateNames(10)
//		fmt.Println(names)
//	}
package namegenerator
