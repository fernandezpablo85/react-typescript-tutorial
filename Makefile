.PHONY: build test run clean-ts

build:
	@go build -v ./...

test:
	@go test ./...

clean-ts:
	rm -rf public/js/compiled

ts_files = $(shell find public/typescript -name '*.tsx')
ts-compile: $(ts_files)
	@tsc --noImplicitAny --jsx react --rootDir public/typescript --outDir public/js/compiled $?

ts-compile-watch: $(ts_files)
	@tsc -w --noImplicitAny --jsx react --rootDir public/typescript --outDir public/js/compiled $?

run: ts-compile
	@go run main.go