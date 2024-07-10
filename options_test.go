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

func TestOptions(t *testing.T) {
	testcase := []struct {
		description         string
		options             []namegenerator.Options
		prefix              string
		suffix              string
		expectedFirstNames  []string
		expectedSecondNames []string
		separator           string
	}{
		{
			description:         "without gender option",
			options:             []namegenerator.Options{},
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "-",
		},
		{
			description:         "with male gender option",
			options:             []namegenerator.Options{namegenerator.WithGender(namegenerator.Male)},
			expectedFirstNames:  namegenerator.MaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			separator:           "-",
		},
		{
			description:         "with female gender option",
			options:             []namegenerator.Options{namegenerator.WithGender(namegenerator.Female)},
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			separator:           "-",
		},
		{
			description:         "with non-binary gender option",
			options:             []namegenerator.Options{namegenerator.WithGender(namegenerator.NonBinary)},
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "-",
		},
		{
			description:         "without prefix option",
			options:             []namegenerator.Options{},
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "-",
		},
		{
			description:         "with prefix option",
			options:             []namegenerator.Options{namegenerator.WithPrefix("Dr. ")},
			prefix:              "Dr. ",
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "-",
		},
		{
			description:         "without suffix option",
			options:             []namegenerator.Options{},
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "-",
		},
		{
			description:         "with suffix option",
			options:             []namegenerator.Options{namegenerator.WithSuffix(" Jr.")},
			suffix:              " Jr.",
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "-",
		},
		{
			description:         "with prefix and suffix option",
			options:             []namegenerator.Options{namegenerator.WithPrefix("Dr. "), namegenerator.WithSuffix(" Jr.")},
			prefix:              "Dr. ",
			suffix:              " Jr.",
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "-",
		},
		{
			description:         "with female gender and prefix and suffix option",
			options:             []namegenerator.Options{namegenerator.WithGender(namegenerator.Female), namegenerator.WithPrefix("Dr. "), namegenerator.WithSuffix(" Jr.")},
			prefix:              "Dr. ",
			suffix:              " Jr.",
			expectedFirstNames:  namegenerator.FemaleNames,
			expectedSecondNames: namegenerator.FamilyNames,
			separator:           "-",
		},
		{
			description:         "with random string option",
			options:             []namegenerator.Options{namegenerator.WithRandomString(5)},
			suffix:              "",
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "-",
		},
		{
			description:         "with default separator",
			options:             []namegenerator.Options{namegenerator.WithSeparator("-")},
			suffix:              "",
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "-",
		},
		{
			description:         "with different separator",
			options:             []namegenerator.Options{namegenerator.WithSeparator("*")},
			suffix:              "",
			expectedFirstNames:  namegenerator.GeneralNames,
			expectedSecondNames: namegenerator.GeneralNames,
			separator:           "*",
		},
	}
	for _, tc := range testcase {
		t.Run(tc.description, func(t *testing.T) {
			nameGenerator := namegenerator.NewGenerator()

			name := nameGenerator.Generate(tc.options...)
			test(t, name, tc.prefix, tc.suffix, tc.separator, tc.expectedFirstNames, tc.expectedSecondNames)

			names := nameGenerator.GenerateMultiple(5, tc.options...)

			if len(names) != 5 {
				t.Errorf("Expected 5 names, got %d", len(names))
			}

			for _, name := range names {
				test(t, name, tc.prefix, tc.suffix, tc.separator, tc.expectedFirstNames, tc.expectedSecondNames)
			}
		})
	}
}
