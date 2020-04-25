export GO111MODULE = on

.PHONY: default test test-cover dev

lint:
	golangci-lint run

# for test
test:
	go test -race -cover ./...

test-cover:
	go test -race -coverprofile=test.out ./... && go tool cover --html=test.out
	