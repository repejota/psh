VERSION=`cat VERSION`
BUILD=`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:%h -1`
PACKAGES = "./..."

# Setup the -ldflags option for go build here, interpolate the variable
# values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

install:
	go install $(LDFLAGS) -v $(PACKAGES)

build:
	go build $(LDFLAGS) -v $(PACKAGES)

# Testing

.PHONY: test
test:
	go test -v $(PACKAGES)

.PHONY: cover
cover:
	go test -cover $(PACKAGES)

.PHONY: cover-html
cover-html:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	rm -rf coverage.out
	go tool cover -html=coverage-all.out

# Lint

lint:
	gometalinter --tests ./...

deps:
	go get -u github.com/spf13/cobra/cobra

dev-deps:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

dist: dist-linux dist-darwin dist-windows

dist-linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o psh-${VERSION}-linux-amd64
	zip psh-${VERSION}-linux-amd64.zip psh-${VERSION}-linux-amd64 README.md LICENSE
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o psh-${VERSION}-linux-386
	zip psh-${VERSION}-linux-386.zip psh-${VERSION}-linux-386 README.md LICENSE

dist-darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o psh-${VERSION}-darwin-amd64
	zip psh-${VERSION}-darwin-amd64.zip psh-${VERSION}-darwin-amd64 README.md LICENSE
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o psh-${VERSION}-darwin-386
	zip psh-${VERSION}-darwin-386.zip psh-${VERSION}-darwin-386 README.md LICENSE

dist-windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o psh-${VERSION}-windows-amd64.exe
	zip psh-${VERSION}-windows-amd64.zip psh-${VERSION}-windows-amd64.exe README.md LICENSE
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o psh-${VERSION}-windows-386.exe
	zip psh-${VERSION}-windows-386.zip psh-${VERSION}-windows-386.exe README.md LICENSE

clean:
	go clean
	rm -rf psh-*
	rm -rf cover.out

reload:
	source `pwd`/contrib/install.sh

