# Application name
APP_NAME = todo-cli

# Output directory
OUTPUT_DIR = bin

# Supported platforms
PLATFORMS = windows/amd64 linux/amd64 darwin/amd64 darwin/arm64

# Build task
.PHONY: build
build: clean
	mkdir -p $(OUTPUT_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d/ -f1); \
		GOARCH=$$(echo $$platform | cut -d/ -f2); \
		OUTPUT_NAME=$(OUTPUT_DIR)/$(APP_NAME)-$$GOOS-$$GOARCH; \
		if [ $$GOOS = "windows" ]; then OUTPUT_NAME+=".exe"; fi; \
		echo "Building $$OUTPUT_NAME..."; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$OUTPUT_NAME; \
	done
	@echo "Build complete! Executables are in the $(OUTPUT_DIR) directory."

# Clean task
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(OUTPUT_DIR)

# Run task for local development
.PHONY: run
run:
	go run main.go

# Test task
.PHONY: test
test:
	go test ./... -v
