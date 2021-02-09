BIN_DIR ?= .
VERSION ?= $(shell git describe --tags --dirty --exact-match 2>/dev/null || git rev-parse --short HEAD)
GO_LDFLAGS = -ldflags "-X main.Version=$(VERSION)"

P = ${BIN_DIR}/gpsdclient
build:
	GOOS=linux GOARCH=arm64 go build $(GO_LDFLAGS) -o ${P} main.go

clean:
	rm -f ${P}

.PHONY: clean
