BINARY_NAME=gocn-server
BINARY_PATH=./cmd/gocn-server/main.go

.PHONY: all build clean test run docker-build docker-up docker-down docker-logs

all: build

# Go tools
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o bin/$(BINARY_NAME) $(BINARY_PATH)

clean:
	@echo "Cleaning..."
	@go clean
	@rm -f bin/$(BINARY_NAME)

test:
	@echo "Running tests..."
	@go test -v ./...

run: build
	@echo "Running $(BINARY_NAME)..."
	@./bin/$(BINARY_NAME)

# Docker commands
docker-build:
	@echo "Building Docker image for $(BINARY_NAME)..."
	@docker-compose -f ./deployments/docker-compose.yml build

docker-up:
	@echo "Starting Docker services..."
	@docker-compose -f ./deployments/docker-compose.yml up
	
docker-down:
	@echo "Stopping Docker services..."
	@docker-compose -f ./deployments/docker-compose.yml down

docker-logs:
	@echo "Showing Docker logs..."
	@docker-compose -f ./deployments/docker-compose.yml logs -f

# Create a bin directory if it doesn't exist
$(shell mkdir -p bin) 