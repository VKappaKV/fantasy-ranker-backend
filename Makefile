.PHONY: run test tidy lint

run:
	air

test:
	go test ./... -count=1

tidy:
	go mod tidy
