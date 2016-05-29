##########
default: build
##########

build: deps lint test
	gox -osarch="darwin/amd64" -output=".bin/{{.OS}}-{{.Arch}}/golang-vuejs-starter-kit" -verbose
	# gox -output=".bin/{{.OS}}-{{.Arch}}/goLang-api" -verbose
	# gox -osarch="linux/amd64 darwin/amd64" -output=".bin/{{.OS}}-{{.Arch}}/golang-vuejs-starter-kit" -verbose

clean:
	rm -rf ./.bin

deps:
	godep go install -v -x ./...

errcheck:
	# errcheck github.com/adam-hanna/goLang-api

fmt:
	gofmt -e -s -w .

lint: fmt errcheck

test: lint
	# TODO(@adam-hanna): Add some damn tests to fix this.

.PHONY: default
.PHONY: build clean deps errcheck fmt lint test