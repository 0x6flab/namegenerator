.PHONY: generate
generate:
	python namegenerator.py
	gofmt -w names.go

.PHONY: lint
lint:
	golangci-lint cache clean && golangci-lint run

.PHONY: examples
examples:
	go run examples/female/main.go && \
	go run examples/male/main.go && \
	go run examples/multiple/main.go && \
	go run examples/nonbinary/main.go

.PHONY: test
test:
	go test --race -covermode=atomic -coverprofile cover.out ./...

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
