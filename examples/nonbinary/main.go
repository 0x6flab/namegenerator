// Copyright (c) 0x6flab. All rights reserved.
//
// SPDX-License-Identifier: GNU GENERAL PUBLIC LICENSE

package main

import (
	"fmt"

	"github.com/0x6flab/namegenerator"
)

func main() {
	// Either of these will work.
	generator := namegenerator.NewNameGenerator()
	name := generator.Generate()
	fmt.Println(name)

	generator = namegenerator.NewNameGenerator().WithGender(namegenerator.NonBinary)
	name = generator.Generate()
	fmt.Println(name)
}
