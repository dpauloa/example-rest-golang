test:
	@go test -count=1 -cover ./...

build:
	@go get ./...

run: build
	@server
