BINARY=psh
VERSION=`cat VERSION`
BUILD=`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:%h -1`
PACKAGES = $(shell go list ./...)

# Setup the -ldflags option for go build here, interpolate the variable
# values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Build & Install

install:
	go install $(LDFLAGS) -v $(PACKAGES)

.PHONY: version
version:
	@echo $(VERSION)-$(BUILD)

# Testing

.PHONY: test
test:
	go test -v $(PACKAGES)

.PHONY: cover-profile
cover-profile:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	rm -rf coverage.out

.PHONY: cover
cover: cover-profile
	go tool cover -func=coverage-all.out

.PHONY: cover-html
cover-html: cover-profile
	go tool cover -html=coverage-all.out

# Lint

lint:
	gometalinter --tests ./...

# Dependencies

deps:

dev-deps:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

# Pacakge and Distribution

dist: dist-linux dist-darwin dist-windows

dist-linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-linux-amd64
	zip ${BINARY}-${VERSION}-linux-amd64.zip ${BINARY}-${VERSION}-linux-amd64 README.md LICENSE
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-linux-386
	zip ${BINARY}-${VERSION}-linux-386.zip ${BINARY}-${VERSION}-linux-386 README.md LICENSE

dist-darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-darwin-amd64
	zip ${BINARY}-${VERSION}-darwin-amd64.zip ${BINARY}-${VERSION}-darwin-amd64 README.md LICENSE
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-darwin-386
	zip ${BINARY}-${VERSION}-darwin-386.zip ${BINARY}-${VERSION}-darwin-386 README.md LICENSE

dist-windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-windows-amd64.exe
	zip ${BINARY}-${VERSION}-windows-amd64.zip ${BINARY}-${VERSION}-windows-amd64.exe README.md LICENSE
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-windows-386.exe
	zip ${BINARY}-${VERSION}-windows-386.zip ${BINARY}-${VERSION}-windows-386.exe README.md LICENSE

# Cleaning up

.PHONY: clean
clean:
	go clean
	rm -rf ${BINARY}
	rm -rf coverage-all.out
	rm -rf ${BINARY}-*

# Run

colortest: install
	psh --colortest

bgtest: install
	psh --backgroundtest

# Docs

godoc-serve:
	godoc -http=":9090"
