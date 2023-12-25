## test: run all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out


## coverage: displays test coverage
coverage:
	@go test -cover ./...
	
## build_cli: builds the command line tool scotch and copies it to scotch-app
build_cli:
	@go build -o ../myapp/scotch ./cmd/cli