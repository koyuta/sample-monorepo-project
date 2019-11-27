GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLIST=$(GOCMD) list
GOTOOL=$(GOCMD) tool

STATIK=statik

VENDOR_DIR=./vendor
CMD_DIR=./cmd
BIN_DIR=./bin

COVERAGE_OUT=cover.out
COVERAGE_HTML=cover.html

#.PHONY: all
#all: dep build-all
#
#.PHONY: install-dep
#install-dep:
#ifeq ($(shell type dep),)
#	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
#endif
#
#.PHONY: dep
#dep: install-dep
#	dep ensure
#
#.PHONY: build
#build:
#	CGO_ENABLED=0 $(GOBUILD) -o $(BIN_DIR)/${BUILD_TARGET} -v $(CMD_DIR)/${BUILD_TARGET}

.PHONY: build
build-all:
	@for c in $$(ls $(CMD_DIR)); do\
		CGO_ENABLED=0 $(GOBUILD) -o $(BIN_DIR)/$$c -v $(CMD_DIR)/$$c;\
	done

.PHONY: test
test:
	$(GOTEST) -coverprofile=$(COVERAGE_OUT) $$($(GOLIST) ./...)
	$(GOTOOL) cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf $(VENDOR_DIR)
	rm -f $(COVERAGE_OUT)
	rm -f $(COVERAGE_HTML)
	rm -f $(BIN_DIR)/*
