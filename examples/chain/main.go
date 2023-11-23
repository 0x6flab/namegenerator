// Copyright (c) 0x6flab. All rights reserved.
//
// SPDX-License-Identifier: GNU GENERAL PUBLIC LICENSE

package main

import (
	"fmt"

	"github.com/0x6flab/namegenerator"
)

func main() {
	generator := namegenerator.NewNameGenerator().WithGender(namegenerator.Male).WithPrefix("Dr. ").WithSuffix(" Jr.")
	name := generator.GenerateNames(10)
	fmt.Println(name)
}
