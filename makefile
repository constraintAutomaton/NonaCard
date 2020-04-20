GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=NoncaCard
BINARY_UNIX=$(BINARY_NAME)_unix
NPM = npm
FRESH = fresh
SASS = sass --watch
CSS_DIR = assets/css
SASS_DIR = assets/sass
TEST_PATH= ./test/...
BUILD_PATH=./build
GOGENERATE=${GOCMD} generate

.PHONY: all test clear

all: test build
build: 
	make submodules
	make run-frontEnd
	${GOCMD} build
test: 
	${GOGENERATE} ./...
	$(GOTEST) -v ${TEST_PATH}
build:
	${GOGET}
	mkdir -p ${BUILD_PATH}
	${GOBUILD} -o ${BUILD_PATH}/${BINARY_NAME}
clear: 
	rm -f ${BUILD_PATH}/$(BINARY_NAME)
	rm -f ${BUILD_PATH}/$(BINARY_UNIX)
run:
	make submodules
	make run-frontEnd
	mkdir -p ${BUILD_PATH}
	build
	${BUILD_PATH}/$(BINARY_NAME)
init:
	build
	cd assets && ${NPM} i
dev:
	run-frontEnd
	${FRESH}
run-frontEnd:
	cd assets && sudo ${NPM} i && ${NPM} start
dev-frontEnd:
	cd assets && sudo ${NPM} i && ${NPM} run-script dev 
submodules:
	git submodule update --init --recursive
lint:
	go gmt ./...
	fgt golint ./...

            