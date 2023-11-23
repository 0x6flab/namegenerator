// Copyright (c) 0x6flab. All rights reserved.
//
// SPDX-License-Identifier: GNU GENERAL PUBLIC LICENSE
package namegenerator_test

import (
	"strings"
	"testing"

	"github.com/0x6flab/namegenerator"
)

const (
	prefix = "Mr. "
	suffix = "@gmail.com"
)

func TestNameGenerator_Generate(t *testing.T) {
	testcase := []struct {
		description         string
		generator           namegenerator.NameGenerator
		expectedFirstNames  []string
		expectedSecondNames []string
		expectedPrefix      string
		expectedSuffix      string
	}{
		{
			description:         "generate default name",
			generator:           namegenerator.NewNameGenerator(),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
		},
		{
			description:         "generate male name",
			generator:           namegenerator.NewNameGenerator().WithGender(namegenerator.Male),
			expectedFirstNames:  namegenerator.MaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
		},
		{
			description:         "generate female name",
			generator:           namegenerator.NewNameGenerator().WithGender(namegenerator.Female),
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
		},
		{
			description:         "generate non-binary name",
			generator:           namegenerator.NewNameGenerator().WithGender(namegenerator.NonBinary),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
		},
		{
			description:         "generate default name with prefix",
			generator:           namegenerator.NewNameGenerator().WithPrefix(prefix),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      prefix,
			expectedSuffix:      "",
		},
		{
			description:         "generate default name with suffix",
			generator:           namegenerator.NewNameGenerator().WithSuffix(suffix),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      "",
			expectedSuffix:      suffix,
		},
		{
			description:         "generate default name with prefix and suffix",
			generator:           namegenerator.NewNameGenerator().WithPrefix(prefix).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      prefix,
			expectedSuffix:      suffix,
		},
		{
			description:         "generate male name with prefix",
			generator:           namegenerator.NewNameGenerator().WithGender(namegenerator.Male).WithPrefix(prefix),
			expectedFirstNames:  namegenerator.MaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      prefix,
			expectedSuffix:      "",
		},
		{
			description:         "generate male name with suffix",
			generator:           namegenerator.NewNameGenerator().WithGender(namegenerator.Male).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.MaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      "",
			expectedSuffix:      suffix,
		},
		{
			description:         "generate male name with prefix and suffix",
			generator:           namegenerator.NewNameGenerator().WithGender(namegenerator.Male).WithPrefix(prefix).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.MaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      prefix,
			expectedSuffix:      suffix,
		},
		{
			description:         "generate female name with prefix",
			generator:           namegenerator.NewNameGenerator().WithGender(namegenerator.Female).WithPrefix(prefix),
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      prefix,
			expectedSuffix:      "",
		},
		{
			description:         "generate female name with suffix",
			generator:           namegenerator.NewNameGenerator().WithGender(namegenerator.Female).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      "",
			expectedSuffix:      suffix,
		},
		{
			description:         "generate female name with prefix and suffix",
			generator:           namegenerator.NewNameGenerator().WithGender(namegenerator.Female).WithPrefix(prefix).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      prefix,
			expectedSuffix:      suffix,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.description, func(t *testing.T) {
			name := tc.generator.Generate()

			test(t, name, tc.expectedPrefix, tc.expectedSuffix, tc.expectedFirstNames, tc.expectedSecondNames)

			names := tc.generator.GenerateNames(5)

			if len(names) != 5 {
				t.Errorf("Expected 5 names, got %d", len(names))
			}

			for _, name := range names {
				test(t, name, tc.expectedPrefix, tc.expectedSuffix, tc.expectedFirstNames, tc.expectedSecondNames)
			}
		})
	}
}

func test(t *testing.T, name, prefix, suffix string, firstNames, secondNames []string) {
	if name == "" {
		t.Errorf("Expected a name, got an empty string")
	}

	if prefix != "" && !strings.HasPrefix(name, prefix) {
		t.Errorf("Expected a name with a prefix, got %s", name)
	}
	name = strings.TrimPrefix(name, prefix)

	if suffix != "" && !strings.HasSuffix(name, suffix) {
		t.Errorf("Expected a name with a suffix, got %s", name)
	}
	name = strings.TrimSuffix(name, suffix)

	if strings.Contains(name, "-") {
		if !contains(firstNames, strings.Split(name, "-")[0]) {
			t.Errorf("Generated name '%s' is not contained in the list of first names", name)
		}

		if !contains(secondNames, strings.Split(name, "-")[1]) {
			t.Errorf("Generated name '%s' is not contained in the list of second names", name)
		}
	} else {
		t.Errorf("Expected a name does not contain a family name separator, got %s", name)
	}
}

func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}

	return false
}

func BenchmarkNameGenerator_Generate(b *testing.B) {
	generator := namegenerator.NewNameGenerator()
	for i := 0; i < b.N; i++ {
		generator.Generate()
	}
}

func BenchmarkNameGenerator_Generate10Names(b *testing.B) {
	generator := namegenerator.NewNameGenerator()
	for i := 0; i < b.N; i++ {
		generator.GenerateNames(10)
	}
}

func BenchmarkNameGenerator_Generate1KNames(b *testing.B) {
	generator := namegenerator.NewNameGenerator()
	for i := 0; i < b.N; i++ {
		generator.GenerateNames(1000)
	}
}
