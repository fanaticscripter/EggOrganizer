.PHONY: all build dist darwin windows

VERSION := $(shell date +%Y%m%d)-$(shell git rev-parse --short=7 HEAD)

all: build

build:
	protoc --go_out=paths=source_relative:. ei/ei.proto aux/common.proto
	go build -o EggOrganizer

dist:
	echo $(VERSION)
	mkdir -p dist/EggOrganizer
	protoc --go_out=paths=source_relative:. ei/ei.proto aux/common.proto
	$(MAKE) darwin windows
	install -m644 README.txt dist/EggOrganizer/README.txt
	install -m644 config.template.toml dist/EggOrganizer/config.toml
	cd dist && rm -f EggOrganizer-$(VERSION).zip && zip -r EggOrganizer-$(VERSION).zip EggOrganizer

darwin:
	GOOS=darwin GOARCH=amd64 GOFLAGS=-trimpath go build -ldflags="-s -w -X github.com/fanaticscripter/EggOrganizer/cmd.Version=$(VERSION)" -o dist/EggOrganizer/EggOrganizer
	upx dist/EggOrganizer/EggOrganizer

windows:
	GOOS=windows GOARCH=amd64 GOFLAGS=-trimpath go build -ldflags="-s -w -X github.com/fanaticscripter/EggOrganizer/cmd.Version=$(VERSION)" -o dist/EggOrganizer/EggOrganizer.exe
	upx dist/EggOrganizer/EggOrganizer.exe
