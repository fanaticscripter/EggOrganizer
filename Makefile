.PHONY: all build dist darwin windows

all: build

build:
	protoc --go_out=paths=source_relative:. ei/ei.proto aux/common.proto
	go build -o EggOrganizer

dist:
	mkdir -p dist/EggOrganizer
	protoc --go_out=paths=source_relative:. ei/ei.proto aux/common.proto
	$(MAKE) darwin windows
	install -m644 README.txt dist/EggOrganizer/README.txt
	install -m644 config.template.toml dist/EggOrganizer/config.toml

darwin:
	GOOS=darwin GOARCH=amd64 GOFLAGS=-trimpath go build -ldflags="-s -w" -o dist/EggOrganizer/EggOrganizer
	upx dist/EggOrganizer/EggOrganizer

windows:
	GOOS=windows GOARCH=amd64 GOFLAGS=-trimpath go build -ldflags="-s -w" -o dist/EggOrganizer/EggOrganizer.exe
	upx dist/EggOrganizer/EggOrganizer.exe
