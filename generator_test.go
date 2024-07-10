// Copyright (c) 2023 0x6flab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
// http://www.apache.org/licenses/LICENSE-2.0

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
		separator           string
	}{
		{
			description:         "generate default name",
			generator:           namegenerator.NewGenerator(),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
			separator:           "-",
		},
		{
			description:         "generate male name",
			generator:           namegenerator.NewGenerator().WithGender(namegenerator.Male),
			expectedFirstNames:  namegenerator.MaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
			separator:           "-",
		},
		{
			description:         "generate female name",
			generator:           namegenerator.NewGenerator().WithGender(namegenerator.Female),
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
			separator:           "-",
		},
		{
			description:         "generate non-binary name",
			generator:           namegenerator.NewGenerator().WithGender(namegenerator.NonBinary),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
			separator:           "-",
		},
		{
			description:         "generate default name with prefix",
			generator:           namegenerator.NewGenerator().WithPrefix(prefix),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      prefix,
			expectedSuffix:      "",
			separator:           "-",
		},
		{
			description:         "generate default name with suffix",
			generator:           namegenerator.NewGenerator().WithSuffix(suffix),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      "",
			expectedSuffix:      suffix,
			separator:           "-",
		},
		{
			description:         "generate default name with prefix and suffix",
			generator:           namegenerator.NewGenerator().WithPrefix(prefix).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      prefix,
			expectedSuffix:      suffix,
			separator:           "-",
		},
		{
			description:         "generate male name with prefix",
			generator:           namegenerator.NewGenerator().WithGender(namegenerator.Male).WithPrefix(prefix),
			expectedFirstNames:  namegenerator.MaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      prefix,
			expectedSuffix:      "",
			separator:           "-",
		},
		{
			description:         "generate male name with suffix",
			generator:           namegenerator.NewGenerator().WithGender(namegenerator.Male).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.MaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      "",
			expectedSuffix:      suffix,
			separator:           "-",
		},
		{
			description:         "generate male name with prefix and suffix",
			generator:           namegenerator.NewGenerator().WithGender(namegenerator.Male).WithPrefix(prefix).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.MaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      prefix,
			expectedSuffix:      suffix,
			separator:           "-",
		},
		{
			description:         "generate female name with prefix",
			generator:           namegenerator.NewGenerator().WithGender(namegenerator.Female).WithPrefix(prefix),
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      prefix,
			expectedSuffix:      "",
			separator:           "-",
		},
		{
			description:         "generate female name with suffix",
			generator:           namegenerator.NewGenerator().WithGender(namegenerator.Female).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      "",
			expectedSuffix:      suffix,
			separator:           "-",
		},
		{
			description:         "generate female name with prefix and suffix",
			generator:           namegenerator.NewGenerator().WithGender(namegenerator.Female).WithPrefix(prefix).WithSuffix(suffix),
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			expectedPrefix:      prefix,
			expectedSuffix:      suffix,
			separator:           "-",
		},
		{
			description:         "generate name with random string",
			generator:           namegenerator.NewGenerator().WithRandomString(5),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
			separator:           "-",
		},
		{
			description:         "generate name with separator as asterisk",
			generator:           namegenerator.NewGenerator().WithSeparator("*"),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
			separator:           "*",
		},
		{
			description:         "generate name with separator as space",
			generator:           namegenerator.NewGenerator().WithSeparator(" "),
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			expectedPrefix:      "",
			expectedSuffix:      "",
			separator:           " ",
		},
	}

	for _, tc := range testcase {
		t.Run(tc.description, func(t *testing.T) {
			name := tc.generator.Generate()

			test(t, name, tc.expectedPrefix, tc.expectedSuffix, tc.separator, tc.expectedFirstNames, tc.expectedSecondNames)

			names := tc.generator.GenerateMultiple(5)

			if len(names) != 5 {
				t.Errorf("Expected 5 names, got %d", len(names))
			}

			for _, name := range names {
				test(t, name, tc.expectedPrefix, tc.expectedSuffix, tc.separator, tc.expectedFirstNames, tc.expectedSecondNames)
			}
		})
	}
}

func test(t *testing.T, name, prefix, suffix, separator string, firstNames, secondNames []string) {
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

	if !strings.Contains(name, separator) {
		t.Errorf("Expected a name does not contain a family name separator, got %s", name)
	}

	if strings.Count(name, separator) == 1 {
		if !contains(firstNames, strings.Split(name, separator)[0]) {
			t.Errorf("Generated name '%s' is not contained in the list of first names", name)
		}

		if !contains(secondNames, strings.Split(name, separator)[1]) {
			t.Errorf("Generated name '%s' is not contained in the list of second names", name)
		}
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
