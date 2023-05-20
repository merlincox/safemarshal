coverage:
	go test -v -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out

test:
	go test ./...

lint:
	golangci-lint run