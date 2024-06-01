// Copyright (c) 2023 0x6flab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
// http://www.apache.org/licenses/LICENSE-2.0

package namegenerator

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
	"time"
)

func init() {
	mrand.New(mrand.NewSource(time.Now().UnixNano()))
}

type Gender uint8

const (
	Male Gender = iota
	Female
	NonBinary
)

// NameGenerator is an interface for generating names.
type NameGenerator interface {
	// Generate generates a name based on the gender.
	//
	// Example:
	//  generator := namegenerator.NewGenerator()
	//  name := generator.Generate()
	//  fmt.Println(name)
	// Output:
	//  `John-Smith`
	// Example with options:
	//  generator := namegenerator.NewGenerator()
	//  name := generator.Generate(namegenerator.WithGender(namegenerator.Male))
	//  fmt.Println(name)
	// Output:
	//  `John-Smith`
	Generate(options ...Options) string

	// GenerateMultiple generates a list of names.
	//
	// Example:
	//  generator := namegenerator.NewGenerator()
	//  names := generator.GenerateMultiple(10)
	//  fmt.Println(names)
	// Example with options:
	//  generator := namegenerator.NewGenerator()
	//  names := generator.GenerateMultiple(10, namegenerator.WithGender(namegenerator.Male))
	//  fmt.Println(names)
	GenerateMultiple(count int, options ...Options) []string

	// WithGender generates a name based on the gender.
	//
	// Example:
	//  generator := namegenerator.NewGenerator().WithGender(namegenerator.Male)
	//  name := generator.Generate()
	//  fmt.Println(name)
	// Output:
	//  `John-Smith`
	WithGender(gender Gender) NameGenerator

	// WithPrefix generates a name with a prefix.
	//
	// Example:
	//  generator := namegenerator.NewGenerator().WithPrefix("Mr. ")
	//  name := generator.Generate()
	//  fmt.Println(name)
	// Output:
	//  `Mr. John-Smith`
	WithPrefix(prefix string) NameGenerator

	// WithSuffix generates a name with a suffix.
	//
	// Example:
	//  generator := namegenerator.NewGenerator().WithSuffix("@gmail.com")
	//  name := generator.Generate()
	//  fmt.Println(name)
	// Output:
	//  `John-Smith@gmail.com`
	WithSuffix(suffix string) NameGenerator

	// WithRandomString generates a name with a random string.
	//
	// Example:
	//  generator := namegenerator.NewGenerator().WithRandomString(5)
	//  name := generator.Generate()
	//  fmt.Println(name)
	// Output:
	//  `John-Smith-abcde`
	WithRandomString(lenOfRandomString int) NameGenerator
}

// nameGenerator is a struct that implements NameGenerator.
type nameGenerator struct {
	gender            Gender
	prefix            string
	suffix            string
	lenOfRandomString int
}

// NewGenerator returns a new NameGenerator.
//
// Example to generate general names:
//
//	generator := namegenerator.NewGenerator()
//
// Example to generate male names:
//
//	generator := namegenerator.NewGenerator().WithGender(namegenerator.Male)
//
// Example to generate female names:
//
//	generator := namegenerator.NewGenerator().WithGender(namegenerator.Female)
//
// Example to generate non-binary names:
//
//	generator := namegenerator.NewGenerator().WithGender(namegenerator.NonBinary)
//
// Example to generate names with a prefix:
//
//	generator := namegenerator.NewGenerator().WithPrefix("Mr. ")
//
// Example to generate names with a suffix:
//
//	generator := namegenerator.NewGenerator().WithSuffix("@gmail.com")
//
// Example to generate names with a random string:
//
//	generator := namegenerator.NewGenerator().WithRandomString(5)
func NewGenerator() NameGenerator {
	return &nameGenerator{
		gender:            NonBinary,
		prefix:            "",
		suffix:            "",
		lenOfRandomString: 0,
	}
}

func (namegen *nameGenerator) WithGender(gender Gender) NameGenerator {
	namegen.gender = gender

	return namegen
}

func (namegen *nameGenerator) WithPrefix(prefix string) NameGenerator {
	namegen.prefix = prefix

	return namegen
}

func (namegen *nameGenerator) WithSuffix(suffix string) NameGenerator {
	namegen.suffix = suffix

	return namegen
}

func (namegen *nameGenerator) WithRandomString(lenOfRandomString int) NameGenerator {
	namegen.lenOfRandomString = lenOfRandomString

	return namegen
}

func (namegen *nameGenerator) Generate(options ...Options) string {
	for _, option := range options {
		option(namegen)
	}

	randomName := ""
	if namegen.lenOfRandomString > 0 {
		randomName = "-" + namegen.generateRandomString(namegen.lenOfRandomString)
	}

	switch namegen.gender {
	case Male:
		grandom := namegen.generateRandomNumber(len(MaleNames))
		frandom := namegen.generateRandomNumber(len(FamilyNames))

		return namegen.prefix + MaleNames[grandom] + "-" + FamilyNames[frandom] + randomName + namegen.suffix
	case Female:
		grandom := namegen.generateRandomNumber(len(FemaleNames))
		frandom := namegen.generateRandomNumber(len(FamilyNames))

		return namegen.prefix + FemaleNames[grandom] + "-" + FamilyNames[frandom] + randomName + namegen.suffix
	case NonBinary:
		fallthrough
	// This condition never happens, but it's here to make the compiler happy.
	default:
		g1random := namegen.generateRandomNumber(len(GeneralNames))
		g2random := namegen.generateRandomNumber(len(GeneralNames))

		return namegen.prefix + GeneralNames[g1random] + "-" + GeneralNames[g2random] + randomName + namegen.suffix
	}
}

func (namegen *nameGenerator) GenerateMultiple(count int, options ...Options) []string {
	for _, option := range options {
		option(namegen)
	}

	names := make([]string, count)
	for i := 0; i < count; i++ {
		names[i] = namegen.Generate()
	}

	return names
}

// generateRandomNumber generates a random number.
// If the random number generator fails, it will use the math/rand package.
func (namegen *nameGenerator) generateRandomNumber(max int) uint64 {
	random, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		random = big.NewInt(mrand.Int63n(int64(max)))
	}

	return random.Uint64()
}

// generateRandomString generates a random string.
func (namegen *nameGenerator) generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		random := namegen.generateRandomNumber(len(charset))
		b[i] = charset[random]
	}

	return string(b)
}
