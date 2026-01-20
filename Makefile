.PHONY: run test tidy lint

run:
	go run ./cmd/api

test:
	go test ./... -count=1

tidy:
	go mod tidy
