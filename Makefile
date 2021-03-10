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

release:
	# goreleaser --snapshot --skip-publish
	rm -rf $(DIST_FOLDER)
	git tag -a v$(version) -m "Release"
	git push origin v$(version)
	goreleaser --snapshot --rm-dist
	goreleaser release --rm-dist

dryrun:
	#goreleaser --snapshot --skip-publish --rm-dist
	goreleaser build
