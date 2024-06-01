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
	testCases := []struct {
		description string
		options     []namegenerator.Options
	}{
		{
			description: "No options",
			options:     []namegenerator.Options{},
		},
		{
			description: "With Male Gender",
			options:     []namegenerator.Options{namegenerator.WithGender(namegenerator.Male)},
		},
		{
			description: "With Female Gender",
			options:     []namegenerator.Options{namegenerator.WithGender(namegenerator.Female)},
		},
		{
			description: "With Random String of length 2",
			options:     []namegenerator.Options{namegenerator.WithRandomString(2)},
		},
		{
			description: "With Random String of length 5",
			options:     []namegenerator.Options{namegenerator.WithRandomString(5)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			generator := namegenerator.NewGenerator()

			metrics(t, generator, tc.options, tc.description, 1000)
			metrics(t, generator, tc.options, tc.description, 10000)
			metrics(t, generator, tc.options, tc.description, 100000)
			// metrics(t, generator, tc.options, tc.description, 1000000)
		})
	}
}

func getNonUnique(names []string) int {
	unique := 0
	uniqueNames := make(map[string]bool)
	for _, name := range names {
		if _, ok := uniqueNames[name]; !ok {
			unique++
			uniqueNames[name] = true
		}
	}

	return len(names) - unique
}

func metrics(t *testing.T, generator namegenerator.NameGenerator, options []namegenerator.Options, desc string, num int) {
	max := 0
	min := num
	average := 0
	iter := 30
	if num == 1000000 {
		iter = 11
	}
	for i := 0; i < iter; i++ {
		names := generator.GenerateMultiple(num, options...)
		nonUnique := getNonUnique(names)
		if nonUnique > max {
			max = nonUnique
		}
		if nonUnique < min {
			min = nonUnique
		}
		average += nonUnique
	}
	average /= iter

	t.Logf("%d %s: Max %d, Min %d, Average %d", num, desc, max, min, average)
}
