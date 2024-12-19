PROJECT ?= github.com/pu4mane/go-docker-k8s-demo
APP ?= demo
RELEASE ?= 0.0.1
COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GOOS ?= darwin
GOARCH ?= amd64

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w -X ${PROJECT}/pkg/version.Release=${RELEASE} -X ${PROJECT}/pkg/version.Commit=${COMMIT} -X ${PROJECT}/pkg/version.BuildTime=${BUILD_TIME}" -o ${APP} ./cmd/main.go

run: build
	./${APP} -config=./configs/config.yaml