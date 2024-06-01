// Copyright (c) 2023 0x6flab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
// http://www.apache.org/licenses/LICENSE-2.0

package namegenerator

type Options func(*nameGenerator)

// WithGender generates a name based on the gender.
func WithGender(gender Gender) Options {
	return func(n *nameGenerator) {
		n.gender = gender
	}
}

// WithPrefix generates a name with a prefix.
func WithPrefix(prefix string) Options {
	return func(n *nameGenerator) {
		n.prefix = prefix
	}
}

// WithSuffix generates a name with a suffix.
func WithSuffix(suffix string) Options {
	return func(n *nameGenerator) {
		n.suffix = suffix
	}
}

// WithRandomString generates a name with a random string.
func WithRandomString(length int) Options {
	return func(n *nameGenerator) {
		n.lenOfRandomString = length
	}
}
