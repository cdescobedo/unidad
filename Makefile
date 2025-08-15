.PHONY: test fmt vet

test:
	go test ./...

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

run: vet
	go run .

