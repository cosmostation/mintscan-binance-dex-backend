NAME 			    := mintscan
VERSION 			:= $(shell git describe --tags --abbrev=0)
COMMIT 				:= $(shell git log -1 --format='%H')
DESTDIR             ?= $(GOPATH)/bin/${NAME}
BUILD_FLAGS 		:= -ldflags "-X github.com/cosmostation/mintscan-binance-dex-backend/mintscan/main.Version=${VERSION} \
								  -X github.com/cosmostation/mintscan-binance-dex-backend/mintscan/main.Commit=${COMMIT}"

## Show all make target commands.
help:
	@make2help $(MAKEFILE_LIST)

## Print out application information.
version: 
	@echo "NAME: ${NAME}"
	@echo "VERSION: ${VERSION}"
	@echo "COMMIT: ${COMMIT}"	

## Build executable file in $DESTDIR  directory.
build: go.sum
	@echo "-> Building ${NAME} executable file..."
	@go build -mod=readonly $(BUILD_FLAGS) -o $(DESTDIR) .

## Install executable file in $GOBIN direcotry. 
install: go.sum
	@echo "-> Installing ${NAME} executable file..."
	@go install -mod=readonly $(BUILD_FLAGS) .

## Create systemd service file in /etc/systemd/system/{SERVICE_NAME}
make_service:
	@echo "-> Creating ${NAME}.service file..."
	@bash ./service.template

## Move systemd service file and enable it.
enable_service:
	@echo "-> Moving ${NAME} service to /etc/systemd/system/..."
	$(shell sudo mv /home/ubuntu/${NAME}.service /etc/systemd/system/${NAME}.service)
	@echo "-> Enabling ${NAME}.service..."
	$(shell sudo systemctl enable ${NAME}.service)

## Clean the executable file
clean:
	@echo "-> Cleaning ${NAME} binary..."
	rm -f $(DESTDIR) 2> /dev/null

.PHONY: help version build install make_service enable_service clean