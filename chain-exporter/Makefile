VERSION               := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT                := $(shell git log -1 --format='%H')
DESTDIR         	  ?= $(GOPATH)/bin/chain-exporter
BUILD_FLAGS 		  := -ldflags "-X github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/exporter.Version=${VERSION} \
									-X github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/exporter.Commit=${COMMIT}"

# Build executable file in $DESTDIR directory.
build: go.sum
	@echo "-> Building chain-exporter binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o $(DESTDIR) .

# Install executable file in $GOBIN direcotry. 
install: go.sum
	@echo "-> Installing chain-exporter binary..."
	@go install -mod=readonly $(BUILD_FLAGS) .

# Create systemd service file using service template.
make_service:
	@echo "-> Creating chain-exporter.service file..."
	@bash ./service.template

# Move systemd service file and enable it.
enable_service:
	@echo "-> Moving chain-exporter service to /etc/systemd/system/..."
	$(shell sudo mv /home/ubuntu/chain-exporter.service /etc/systemd/system/chain-exporter.service)
	@echo "-> Enabling chain-exporter.service..."
	$(shell sudo systemctl enable chain-exporter.service)

# Clean up all generated executable files.
clean:
	@echo "-> Cleaning chain-exporter binary..."
	rm -f $(TOOLS_DESTDIR) 2> /dev/null

PHONY: build install make_service enable_service clean