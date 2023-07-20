include .env
export

sqlc:
	sqlc generate
.PHONY: sqlc

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci