package namegenerator_test

import (
	"strings"
	"testing"

	"github.com/0x6flab/namegenerator"
)

func TestNewNameGenerator(t *testing.T) {
	generator := namegenerator.NewNameGenerator("")
	if generator == nil {
		t.Errorf("Expected NameGenerator, got nil")
	}
}

func TestNameGenerator_Generate(t *testing.T) {
	generator := namegenerator.NewNameGenerator("")
	name := generator.Generate()

	if len(name) == 0 {
		t.Errorf("Expected a name, got an empty string")
	}

	if !strings.Contains(name, "-") {
		t.Errorf("Expected a name with a family name separator, got %s", name)
	}

	if !contains(namegenerator.GeneralNames, strings.Split(name, "-")[0]) {
		t.Errorf("Generated name '%s' does not contain a general name", name)
	}

	if !contains(namegenerator.FamilyNames, strings.Split(name, "-")[1]) {
		t.Errorf("Generated name '%s' does not contain a family name", name)
	}

	generator = namegenerator.NewNameGenerator("male")
	name = generator.Generate()

	if len(name) == 0 {
		t.Errorf("Expected a male name, got an empty string")
	}

	if !strings.Contains(name, "-") {
		t.Errorf("Expected a male name with a family name separator, got %s", name)
	}

	if !contains(namegenerator.MaleNames, strings.Split(name, "-")[0]) {
		t.Errorf("Generated name '%s' does not contain a male name", name)
	}

	if !contains(namegenerator.FamilyNames, strings.Split(name, "-")[1]) {
		t.Errorf("Generated name '%s' does not contain a family name", name)
	}

	generator = namegenerator.NewNameGenerator("female")
	name = generator.Generate()

	if len(name) == 0 {
		t.Errorf("Expected a female name, got an empty string")
	}

	if !strings.Contains(name, "-") {
		t.Errorf("Expected a female name with a family name separator, got %s", name)
	}

	if !contains(namegenerator.FemaleNames, strings.Split(name, "-")[0]) {
		t.Errorf("Generated name '%s' does not contain a female name", name)
	}

	if !contains(namegenerator.FamilyNames, strings.Split(name, "-")[1]) {
		t.Errorf("Generated name '%s' does not contain a family name", name)
	}
}

func TestNameGenerator_GenerateNames(t *testing.T) {
	generator := namegenerator.NewNameGenerator("")
	names := generator.GenerateNames(5)

	if len(names) != 5 {
		t.Errorf("Expected 5 names, got %d", len(names))
	}

	for _, name := range names {
		if len(name) == 0 {
			t.Errorf("Expected a name, got an empty string")
		}

		if !strings.Contains(name, "-") {
			t.Errorf("Expected a name with a family name separator, got %s", name)
		}

		if !contains(namegenerator.GeneralNames, strings.Split(name, "-")[0]) {
			t.Errorf("Generated name '%s' does not contain a general name", name)
		}

		if !contains(namegenerator.FamilyNames, strings.Split(name, "-")[1]) {
			t.Errorf("Generated name '%s' does not contain a family name", name)
		}
	}

	generator = namegenerator.NewNameGenerator("male")
	names = generator.GenerateNames(5)

	if len(names) != 5 {
		t.Errorf("Expected 5 male names, got %d", len(names))
	}

	for _, name := range names {
		if len(name) == 0 {
			t.Errorf("Expected a male name, got an empty string")
		}

		if !strings.Contains(name, "-") {
			t.Errorf("Expected a male name with a family name separator, got %s", name)
		}

		if !contains(namegenerator.MaleNames, strings.Split(name, "-")[0]) {
			t.Errorf("Generated name '%s' does not contain a male name", name)
		}

		if !contains(namegenerator.FamilyNames, strings.Split(name, "-")[1]) {
			t.Errorf("Generated name '%s' does not contain a family name", name)
		}
	}

	generator = namegenerator.NewNameGenerator("female")
	names = generator.GenerateNames(5)

	if len(names) != 5 {
		t.Errorf("Expected 5 female names, got %d", len(names))
	}

	for _, name := range names {
		if len(name) == 0 {
			t.Errorf("Expected a male name, got an empty string")
		}

		if !strings.Contains(name, "-") {
			t.Errorf("Expected a male name with a family name separator, got %s", name)
		}

		if !contains(namegenerator.FemaleNames, strings.Split(name, "-")[0]) {
			t.Errorf("Generated name '%s' does not contain a female name", name)
		}

		if !contains(namegenerator.FamilyNames, strings.Split(name, "-")[1]) {
			t.Errorf("Generated name '%s' does not contain a family name", name)
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
