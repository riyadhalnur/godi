.PHONY: test coverage

test:
	@go test -v -cover -coverprofile=cover.out ./...

coverage:
	@go tool cover -func=cover.out

vet:
	@go vet ./...

build:
	@go build $(PWD)/cmd/api

run:
	@go run $(PWD)/cmd/api/main.go
