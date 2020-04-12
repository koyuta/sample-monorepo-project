GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLIST=$(GOCMD) list
GOTOOL=$(GOCMD) tool

STATIK=statik

CMD_DIR=./cmd
BIN_DIR=./bin

COVERAGE_OUT=cover.out
COVERAGE_HTML=cover.html

BUILD_TARGET=api

.PHONY: build
build:
	GOPRIVATE=$(REPO) CGO_ENABLED=0 $(GOBUILD) -o $(BIN_DIR)/${BUILD_TARGET} -v $(CMD_DIR)/${BUILD_TARGET}

.PHONY: build-all
build-all:
	@for c in $$(ls $(CMD_DIR)); do\
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BIN_DIR)/$$c -v $(CMD_DIR)/$$c;\
	done

.PHONY: test
test:
	$(GOTEST) -coverprofile=$(COVERAGE_OUT) $$($(GOLIST) ./...)
	$(GOTOOL) cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)

.PHONY: clean
clean:
	rm -f $(COVERAGE_OUT)
	rm -f $(COVERAGE_HTML)
	rm -f $(BIN_DIR)/*
