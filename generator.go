// Copyright (c) 0x6flab. All rights reserved.
//
// SPDX-License-Identifier: GNU GENERAL PUBLIC LICENSE

package namegenerator

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
	"sync"
)

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
	//  generator := namegenerator.NewNameGenerator()
	//  name := generator.Generate()
	//  fmt.Println(name)
	// Output:
	//  `John-Smith`
	Generate() string

	// GenerateNames generates a list of names.
	//
	// Example:
	//  generator := namegenerator.NewNameGenerator()
	//  names := generator.GenerateNames(10)
	//  fmt.Println(names)
	// Output:
	//  `[Dryke-Monroe Scarface-Lesway Shelden-Corsale Marcus-Ivett Victor-Nesrallah Merril-Gulick Leonardo-Lindler Maurits-Lias Rawley-Connor Elvis-Khouderchah]`
	GenerateNames(int) []string

	// WithGender generates a name based on the gender.
	//
	// Example:
	//  generator := namegenerator.NewNameGenerator().WithGender(namegenerator.Male)
	//  name := generator.Generate()
	//  fmt.Println(name)
	// Output:
	//  `John-Smith`
	WithGender(Gender) NameGenerator
}

// nameGenerator is a struct that implements NameGenerator.
type nameGenerator struct {
	gender Gender
}

// NewNameGenerator returns a new NameGenerator.
//
// Example to generate general names:
//
//	generator := namegenerator.NewNameGenerator()
//
// Example to generate male names:
//
//	generator := namegenerator.NewNameGenerator().WithGender(namegenerator.Male)
//
// Example to generate female names:
//
//	generator := namegenerator.NewNameGenerator().WithGender(namegenerator.Female)
//
// Example to generate non-binary names:
//
//	generator := namegenerator.NewNameGenerator().WithGender(namegenerator.NonBinary)
func NewNameGenerator() NameGenerator {
	return &nameGenerator{
		gender: NonBinary,
	}
}

func (namegen *nameGenerator) WithGender(gender Gender) NameGenerator {
	namegen.gender = gender

	return namegen
}

func (namegen *nameGenerator) Generate() string {
	switch namegen.gender {
	case Male:
		grandom := namegen.generateRandomNumber(len(MaleNames))
		frandom := namegen.generateRandomNumber(len(FamilyNames))

		return MaleNames[grandom] + "-" + FamilyNames[frandom]
	case Female:
		grandom := namegen.generateRandomNumber(len(FemaleNames))
		frandom := namegen.generateRandomNumber(len(FamilyNames))

		return FemaleNames[grandom] + "-" + FamilyNames[frandom]
	case NonBinary:
		g1random := namegen.generateRandomNumber(len(GeneralNames))
		g2random := namegen.generateRandomNumber(len(GeneralNames))

		return GeneralNames[g1random] + "-" + GeneralNames[g2random]
	// This condition never happens, but it's here to make the compiler happy.
	default:
		return ""
	}
}

func (namegen *nameGenerator) GenerateNames(count int) []string {
	var waitGroup sync.WaitGroup
	names := make([]string, count)

	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(index int) {
			defer waitGroup.Done()
			names[index] = namegen.Generate()
		}(i)
	}

	waitGroup.Wait()

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
