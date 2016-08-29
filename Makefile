PACKAGES = $(shell go list ./... | grep -v vendor)
VIEWS = $(wildcard views/*.html)

build: install assets
	@go fmt ${PACKAGES}
	@go build ${PACKAGES}
	@go vet ${PACKAGES}

bindata.go: $(VIEWS)
	@go generate ${PACKAGES}
	@touch $@

assets: bindata.go

glide.lock: glide.yaml
	@glide update
	@glide install
	@touch $@

install: glide.lock

test: build
	@go test ${PACKAGES}

clean-ts:
	@rm -rf public/js/compiled

ts_files = $(shell find public/typescript -name '*.tsx')
ts-compile: $(ts_files)
	@tsc --noImplicitAny --jsx react --rootDir public/typescript --outDir public/js/compiled $?

ts-compile-watch: $(ts_files)
	@tsc -w --noImplicitAny --jsx react --rootDir public/typescript --outDir public/js/compiled $?

run: ts-compile assets
	@go run *.go