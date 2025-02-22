test:
		go test ./...
testsum:
		gotestsum ./...
testcov:
		go test -cover ./...

viewtestcov:
		go test -coverprofile=coverage.out ./...
		go tool cover -html=coverage.out

run: build test testcov
		./properfilter --help
build:
		go build -o properfilter ./cmd/main.go
