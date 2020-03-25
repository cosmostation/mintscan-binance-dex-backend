# VERSION               := $(shell echo $(shell git describe --tags) | sed 's/^v//')
VERSION 			  ?= manual
COMMIT                := $(shell git log -1 --format='%H')
TOOLS_DESTDIR         ?= $(GOPATH)/bin/stats-exporter
BUILD_FLAGS 		  := -ldflags "-X main.buildHash=${COMMIT} -X util.BuildVersion=${VERSION}"

build: go.sum
	@echo "building stats-exporter binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o $(TOOLS_DESTDIR) .
	#GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build $(BUILD_FLAGS) -o $(TOOLS_DESTDIR) .

install: go.sum
	@echo "installing stats-exporter binary..."
	@go install -mod=readonly $(BUILD_FLAGS) .

clean:
	@echo "cleaning stats-exporter binary..."
	rm -f $(TOOLS_DESTDIR) 2> /dev/null
