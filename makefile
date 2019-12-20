GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
DEP = dep
BINARY_NAME=3_by_3
BINARY_UNIX=$(BINARY_NAME)_unix
TEST_PATH = ./pkg/test
NPM = npm
FRESH = fresh
SASS = sass --watch
CSS_DIR = assets/css
SASS_DIR = assets/sass
    
all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ${TEST_PATH}
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
deps:
	${DEP} ensure
init:
	${DEP} init
	${NPM} i
dev:
	${FRESH}
dev-css:
	${SASS} ${SASS_DIR}:${CSS_DIR}

            