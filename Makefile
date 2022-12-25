.PHONY: run build build_all build_linux

run:
	wails dev

build:
	wails build

build_all:
	wails build -platform=darwin,windows/amd64,linux