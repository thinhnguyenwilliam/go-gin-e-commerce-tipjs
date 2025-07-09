## how to use
# Command	Action:
# make build	Build the binary
# make run	Run the app using go run
# make tidy	Clean up module imports
# make clean	Remove built files
# make test	Run unit tests
# make fmt	Format code
#make dev	Run with hot-reload via Air
##



# Project Variables
APP_NAME := ecommerce-ver-2
MAIN_FILE := main.go
CMD_DIR := ./cmd/server

# Default target
.PHONY: all
all: build

# Build the app
.PHONY: build
build:
	go build -o bin/$(APP_NAME) $(MAIN_FILE)

# Run the app (assumes main.go in root and config/main.yaml exists)
.PHONY: run
run:
	go run $(CMD_DIR)/$(MAIN_FILE)

# Tidy up modules
.PHONY: tidy
tidy:
	go mod tidy

# Clean built files
.PHONY: clean
clean:
	go clean
	rm -f bin/$(APP_NAME)

# Run tests
.PHONY: test
test:
	go test ./...

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint (if you have golangci-lint or similar installed)
.PHONY: lint
lint:
	golangci-lint run ./...

# Use air for hot-reload development (if installed)
.PHONY: dev
dev:
	air

# Print current working directory (for debug)
.PHONY: pwd
pwd:
	@pwd
