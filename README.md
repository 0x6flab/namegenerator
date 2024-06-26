# namegenerator

[![Coverage Status](https://coveralls.io/repos/github/0x6flab/namegenerator/badge.svg?branch=main)](https://coveralls.io/github/0x6flab/namegenerator?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/0x6flab/namegenerator)](https://goreportcard.com/report/github.com/0x6flab/namegenerator)
[![Go Reference](https://pkg.go.dev/badge/github.com/0x6flab/namegenerator.svg)](https://pkg.go.dev/github.com/0x6flab/namegenerator)
![GitHub](https://img.shields.io/github/license/0x6flab/namegenerator?style=plastic)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/0x6flab/namegenerator?style=plastic)
![GitHub repo size](https://img.shields.io/github/repo-size/0x6flab/namegenerator?style=plastic)
[![Maintainability](https://api.codeclimate.com/v1/badges/d2a9668083e57e08c20b/maintainability)](https://codeclimate.com/github/0x6flab/namegenerator/maintainability)
[![Continuous Integration](https://github.com/0x6flab/namegenerator/actions/workflows/ci.yaml/badge.svg)](https://github.com/0x6flab/namegenerator/actions/workflows/ci.yaml)

NameGenerator is a Golang package that provides a simple yet powerful tool for generating human-readable and pronounceable random names. The names are generated by [namegenerator.py](./namegenerator.py) from [https://www.cs.cmu.edu/][names-dataset] dataset.

## Installation

```bash
go get github.com/0x6flab/namegenerator
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/0x6flab/namegenerator"
)

func main() {
    generator := namegenerator.NewGenerator()

    // Generate a random name
    name := generator.Generate()

    // Print the name
    fmt.Println(name)
}
```

## Examples

See [examples](./examples) directory for more examples.

## Benchmarks

Benchmark tests can be found in [test](benchmark_test.go) file and collision tests can be found in [test](collision_test.go) file.

```bash
goos: linux
goarch: amd64
pkg: github.com/0x6flab/namegenerator
cpu: AMD Ryzen 7 7735HS with Radeon Graphics
BenchmarkNameGenerator_Generate-16                967677              1158 ns/op
BenchmarkNameGenerator_Generate10Names-16          97788             11577 ns/op
BenchmarkNameGenerator_Generate1KNames-16            973           1156777 ns/op
BenchmarkNameGenerator_Generate10KNames-16            88          11594857 ns/op
```

For collision tests:

| Number of Names | Options                       | Collisions | Collision range |
| --------------- | ----------------------------- | ---------- | --------------- |
| 1K              | with random string (2) or (5) | 0          | 0               |
| 10K             | with random string (2) or (5) | 0          | 0               |
| 100K            | with random string (2) or (5) | 0          | 0               |
| 1M              | with random string (2) or (5) | 0          | 0               |
| 1K              | without                       | 0          | 0               |
| 10K             | without                       | 0          | 0               |
| 100K            | without                       | 1          | 0-5             |
| 1M              | without                       | 147        | 135-163         |
| 1K              | male names                    | 0          | 0               |
| 10K             | male names                    | 1          | 0-4             |
| 100K            | male names                    | 127        | 105-152         |
| 1M              | male names                    | 12516      | 12394-12630     |
| 1K              | female names                  | 0          | 0               |
| 10K             | female names                  | 0          | 0-3             |
| 100K            | female names                  | 74         | 57-98           |
| 1M              | female names                  | 7404       | 7289-7542       |

## License

This project is licensed under the Apache-2.0 LICENSE - see the [LICENSE](./LICENSE) file for details.

[names-dataset]: https://www.cs.cmu.edu/afs/cs/project/ai-repository/ai/areas/nlp/corpora/0.html
