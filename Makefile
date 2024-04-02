# Copyright (c) 2023 0x6flab
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at:
# http://www.apache.org/licenses/LICENSE-2.0

.PHONY: generate
generate:
	python namegenerator.py
	gofmt -w names.go

.PHONY: lint
lint:
	golangci-lint cache clean && golangci-lint run

.PHONY: examples
examples:
	go run examples/main.go

.PHONY: test
test:
	go test --race -covermode=atomic -coverprofile cover.out $(go list ./... | grep -v /examples)

.PHONY: bench
bench:
	go test -bench='.' -cpuprofile='cpu.prof' -memprofile='mem.prof'

.PHONY: cover-html
cover-html: test
	go tool cover -html cover.out -o cover.html

godoc-serve:
	godoc -http=:6060

.PHONY: install-precommit
install-precommit:
	pre-commit install

.PHONY: precommit
precommit:
	pre-commit run --all-files -v

.PHONY: update-precommit
update-precommit:
	pre-commit autoupdate

release:
	standard-version
	goreleaser release --clean --release-notes CHANGELOG.md
	git push --follow-tags origin main
