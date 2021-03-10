include .env
export

PROJECTNAME := $(shell basename "$(PWD)")
BINARY_NAME=blueprint
HASH := $(shell git rev-parse --short HEAD)
DIST_FOLDER=./dist

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

run:
	go run main.go

build: clean compile

clean:
	rm -rf $(DIST_FOLDER)

prod:
	# goreleaser --snapshot --skip-publish
	git tag -a v$(BLUEPRINT_VERSION) -m "Release"
	git push origin v$(BLUEPRINT_VERSION)
	goreleaser --snapshot --rm-dist
	goreleaser release

