audit:
	go mod verify
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest

clean:
	rm -f cover.out

coverage:
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out

format:
	go fmt ./...

lint:
	golangci-lint run --enable-all

test:
	go test -v ./...
	go test -race ./...
	go vet

tidy:
	go mod tidy

.PHONY: audit clean coverage format lint test tidy
