GO111MODULES=on

.PHONY: build
## build: build the application
build:
	@echo "Building..."
	go build -o go-ok github.com/jayapriya90/go-ok

.PHONY: run
## run: build and run the application
run:
	go run github.com/jayapriya90/go-ok

.PHONY: fmt
## fmt: format Go source code
fmt:
	go fmt ./...

.PHONY: lint
## lint: lint Go source code
lint:
	golint ./...

.PHONY: help
## help: prints the help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
