VERSION=`cat ./VERSION`

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

install:
	go install -v ${LDFLAGS}

build:
	go build -v ${LDFLAGS}

test:
	go test -v ./... --cover

cover:
	go test -coverprofile cover.out
	goveralls -repotoken k7kIkhp7ub6iz1gIEBGOjmOvhuPJB7Dki

lint:
	golint ./...
	go vet ./...

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
	source ./contrib/install.sh
