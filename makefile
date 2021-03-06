GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=NoncaCard
BINARY_UNIX=$(BINARY_NAME)_unix
TEST_PATH = ./pkg/test
NPM = npm
FRESH = fresh
SASS = sass --watch
CSS_DIR = assets/css
SASS_DIR = assets/sass
    
all: test build
build: 
	make submodules
	make run-frontEnd
	${GOCMD} build
test: 
	$(GOTEST) -v ${TEST_PATH}
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	make submodules
	make run-frontEnd
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
init:
	${GOCMD} build
	cd assets && ${NPM} i
dev:
	make run-frontEnd
	${FRESH}
run-frontEnd:
	cd assets && sudo ${NPM} i && ${NPM} start
dev-frontEnd:
	cd assets && sudo ${NPM} i && ${NPM} run-script dev 
submodules:
	git submodule update --init --recursive
lint:
	fgt golint ./...

            