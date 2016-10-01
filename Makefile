VERSION=0.0.1
TARGETS_NOVENDOR=$(shell glide novendor)

all: build

bundle:
	glide install && glide update

fmt:
	@echo $(TARGETS_NOVENDOR) | xargs go fmt

test: build
	./bin/goflake-grpcserver &
	go test -cover $(TARGETS_NOVENDOR)

build:
	go build -o bin/goflake-grpcserver -ldflags "-X main.version=${VERSION}" .

build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/goflake-grpcserver-${VERSION}/goflake-grpcserver -ldflags "-X main.version=${VERSION}" .
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/goflake-grpcserver-${VERSION}/goflake-grpcserver -ldflags "-X main.version=${VERSION}" .

dist: build-all
	cd bin/linux/amd64 && tar zcvf gflake-grpcserver-linux-amd64-${VERSION}.tar.gz goflake-grpcserver-${VERSION}
	cd bin/darwin/amd64 && tar zcvf gflake-grpcserver-darwin-amd64-${VERSION}.tar.gz goflake-grpcserver-${VERSION}

clean:
	rm -rf bin/* 
