// Copyright (c) 2023 0x6flab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
// http://www.apache.org/licenses/LICENSE-2.0

package namegenerator

type Options func(*nameGenerator)

// WithGender generates a name based on the gender.
//
// Example:
//
//	generator := namegenerator.NewGenerator()
//	name := generator.Generate(namegenerator.WithGender(namegenerator.Male))
func WithGender(gender Gender) Options {
	return func(n *nameGenerator) {
		n.gender = gender
	}
}

// WithPrefix generates a name with a prefix.
//
// Example:
//
//	generator := namegenerator.NewGenerator()
//	name := generator.Generate(namegenerator.WithPrefix("Dr. "))
func WithPrefix(prefix string) Options {
	return func(n *nameGenerator) {
		n.prefix = prefix
	}
}

// WithSuffix generates a name with a suffix.
//
// Example:
//
//	generator := namegenerator.NewGenerator()
//	name := generator.Generate(namegenerator.WithSuffix("@gmail.com"))
func WithSuffix(suffix string) Options {
	return func(n *nameGenerator) {
		n.suffix = suffix
	}
}

// WithRandomString generates a name with a random string.
//
// Example:
//
//	generator := namegenerator.NewGenerator()
//	name := generator.Generate(namegenerator.WithRandomString(5))
func WithRandomString(length int) Options {
	return func(n *nameGenerator) {
		n.lenOfRandomString = length
	}
}

// WithSeparator adds the separator in between the names.
//
// Example:
//
//	generator := namegenerator.NewGenerator()
//	name := generator.Generate(namegenerator.WithSeparator("*"))
func WithSeparator(separator string) Options {
	return func(n *nameGenerator) {
		n.separator = separator
	}
}
