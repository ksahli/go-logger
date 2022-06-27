clean:
	go clean -testcache -cache
	rm coverage.out

format:
	go fmt ./...

verify:
	go fmt ./...
	go vet ./...

test:
	go test ./... -coverprofile coverage.out

install:
	go install

