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
	${GOCMD} build
	cd assets && ${NPM} i
test: 
	$(GOTEST) -v ${TEST_PATH}
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
init:
	${GOCMD} build
	cd assets && ${NPM} i
dev:
	${FRESH}
dev-css:
	${SASS} ${SASS_DIR}:${CSS_DIR}
dev-js:
	cd assets && sudo ${NPM} start

            