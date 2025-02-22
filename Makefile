test:
		go test ./...
testsum:
		gotestsum ./...
testcov:
		go test -coverprofile=coverage.out ./...
		go tool cover -func=coverage.out
viewtestcov:
		go test -coverprofile=coverage.out ./...
		go tool cover -html=coverage.out

run: build test testcov
		./properfilter --help
build:
		go get github.com/fatih/color > /dev/null
		go build -o properfilter ./cmd/main.go >/dev/null

clean:
	rm coverage.out

help: build
	./properfilter --help
