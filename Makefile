.PHONY: install

RUN=go run
BUF=$(RUN) github.com/bufbuild/buf/cmd/buf@v1.28.1

build:
	$(BUF) generate