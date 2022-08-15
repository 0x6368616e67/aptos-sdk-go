GO = go
GIT_COMMIT ?= $(shell git rev-list -1 HEAD)
GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
GIT_TAG    ?= $(shell git describe --tags '--match=v*' --dirty)

BUILD_TAGS = 
GO_FLAGS += -trimpath -tags $(BUILD_TAGS) -buildvcs=false
GOTEST = GODEBUG=cgocheck=0 $(GO) test $(GO_FLAGS) ./... -p 2
GOBUILD =  $(GO) build $(GO_FLAGS)

default: all

## all:                               run  all commands
all: build


## test:                              run unit tests with a 50s timeout
test:
	$(GOTEST) --timeout 50s

build:
	@echo "Building $*"
	$(GOBUILD)
	@echo "Done."