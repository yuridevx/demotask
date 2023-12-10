.PHONY: install

VERSION=v1.28.1
BUF=go run github.com/bufbuild/buf/cmd/buf@$(VERSION)

build:
	cd domain
	$(BUF) build