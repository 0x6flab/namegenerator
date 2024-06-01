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

func BenchmarkNameGenerator_Generate(b *testing.B) {
	generator := namegenerator.NewGenerator()
	for i := 0; i < b.N; i++ {
		generator.Generate()
	}
}

func BenchmarkNameGenerator_Generate10Names(b *testing.B) {
	generator := namegenerator.NewGenerator()
	for i := 0; i < b.N; i++ {
		generator.GenerateMultiple(10)
	}
}

func BenchmarkNameGenerator_Generate1KNames(b *testing.B) {
	generator := namegenerator.NewGenerator()
	for i := 0; i < b.N; i++ {
		generator.GenerateMultiple(1000)
	}
}

func BenchmarkNameGenerator_Generate10KNames(b *testing.B) {
	generator := namegenerator.NewGenerator()
	for i := 0; i < b.N; i++ {
		generator.GenerateMultiple(10000)
	}
}
