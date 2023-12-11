.PHONY: generate build

RUN=go run
BUF=$(RUN) github.com/bufbuild/buf/cmd/buf@v1.28.1

generate:
	$(BUF) generate

build:
	cd game && go build -o game ./cmd
	cd numbers && go build -o numbers ./cmd