// Copyright (c) 2023 0x6flab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
// http://www.apache.org/licenses/LICENSE-2.0

package namegenerator_test

import (
	"testing"

	"github.com/0x6flab/namegenerator"
)

func TestCollision(t *testing.T) {
	generator := namegenerator.NewGenerator()
	names := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		name := generator.Generate()
		if _, ok := names[name]; ok {
			t.Errorf("Generated a duplicate name: %s", name)
		}
		names[name] = true
	}
}

func TestMaleCollision(t *testing.T) {
	generator := namegenerator.NewGenerator().WithGender(namegenerator.Male)
	names := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		name := generator.Generate()
		if _, ok := names[name]; ok {
			t.Errorf("Generated a duplicate name: %s", name)
		}
		names[name] = true
	}
}

func TestFemaleCollision(t *testing.T) {
	generator := namegenerator.NewGenerator().WithGender(namegenerator.Female)
	names := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		name := generator.Generate()
		if _, ok := names[name]; ok {
			t.Errorf("Generated a duplicate name: %s", name)
		}
		names[name] = true
	}
}

func TestWithRandomStringCollision(t *testing.T) {
	generator := namegenerator.NewGenerator().WithRandomString(2)
	names := make(map[string]bool)

	for i := 0; i < 1000000; i++ {
		name := generator.Generate()
		if _, ok := names[name]; ok {
			t.Errorf("Generated a duplicate name: %s", name)
		}
		names[name] = true
	}
}
