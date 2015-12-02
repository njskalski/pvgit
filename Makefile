default: deps
VERSION := $(shell git describe --always --dirty --tags)

deps:
	glide -q -y glide.yaml up

install:
	go install -ldflags "-X main.version=${VERSION}"

.PHONY: deps install
