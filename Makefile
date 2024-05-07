CROSS_TC=arm-kobo-linux-gnueabihf
PWD=$(shell pwd)
XTOOLS=${PWD}/x-tools/arm-kobo-linux-gnueabihf/bin/
.PHONY: build
build:
	PATH=${XTOOLS}:$$PATH CGO_ENABLED=1 GOARCH=arm GOOS=linux CC=${CROSS_TC}-gcc CXX=${CROSS_TC}-g++ go build -o ./build/kobowriter
