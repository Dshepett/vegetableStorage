lint:
	golangci-lint run

build:
	go build -o ./bin/app ./cmd/

run:
	go run ./cmd/