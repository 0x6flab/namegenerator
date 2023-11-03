// Copyright (c) 0x6flab. All rights reserved.
//
// SPDX-License-Identifier: GNU GENERAL PUBLIC LICENSE

package main

import (
	"fmt"

	"github.com/0x6flab/namegenerator"
)

func main() {
	generator := namegenerator.NewNameGenerator().WithGender(namegenerator.Female)
	name := generator.Generate()
	fmt.Println(name)
}
