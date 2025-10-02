.PHONY: test test-unit test-integration test-verbose test-coverage clean build help

# Build the migrate binary
build:
	@echo "Building migrate-v2 binary..."
	@go build -o bin/migrate-v2 main.go
	@echo "Binary built: bin/migrate-v2"

# Run all tests
test: test-unit test-integration

# Run unit tests only
test-unit:
	@echo "Running unit tests..."
	@go test -v ./internal/... ./transformations/... ./utils/...

# Run integration tests for all resources
test-integration:
	@echo "Running integration tests..."
	@for dir in resources/*/integration; do \
		if [ -d "$$dir" ]; then \
			echo "Testing $$dir..."; \
			go test -v ./$$dir/... || exit 1; \
		fi; \
	done

# Run integration test for a specific resource
# Usage: make test-resource RESOURCE=access_application
test-resource:
	@echo "Running integration tests for $(RESOURCE)..."
	@go test -v ./resources/$(RESOURCE)/integration/...

# Run specific test case for a resource
# Usage: make test-case RESOURCE=access_application CASE=basic
test-case:
	@echo "Running test case $(CASE) for $(RESOURCE)..."
	@go test -v -run "TestAccessApplicationMigration/$(CASE)" ./resources/$(RESOURCE)/integration/...

# Run tests with verbose output
test-verbose:
	@echo "Running all tests with verbose output..."
	@go test -v -run . ./...

# Run with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -cover -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmark tests
benchmark:
	@echo "Running benchmark tests..."
	@go test -bench=. -benchmem ./resources/*/integration/...

# Clean build and test artifacts
clean:
	@echo "Cleaning artifacts..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@find . -name "*.debug" -delete
	@find . -name "*.backup" -delete
	@echo "Clean complete"

# Validate all test fixtures
validate-fixtures:
	@echo "Validating test fixtures..."
	@for dir in resources/*/integration/testdata; do \
		if [ -d "$$dir" ]; then \
			echo "Checking $$dir..."; \
			for file in $$dir/v4/*.tf $$dir/v5/*.tf; do \
				if [ -f "$$file" ]; then \
					terraform fmt -check $$file || exit 1; \
				fi \
			done \
		fi \
	done
	@echo "All fixtures are valid"

# Update golden files for all resources (USE WITH CAUTION)
update-golden:
	@echo "WARNING: This will update all expected v5 outputs with current migration results"
	@echo "Press Ctrl+C to cancel, or Enter to continue..."
	@read confirm
	@echo "Updating golden files..."
	@for resource in resources/*/integration; do \
		if [ -d "$$resource" ]; then \
			echo "Updating $$resource..."; \
			cd $$resource && go test -v -run TestUpdateGoldenFiles; \
		fi \
	done

# Run the migration tool
# Usage: make migrate FROM=v4 TO=v5 DIR=./test-configs
migrate: build
	@echo "Running migration from $(FROM) to $(TO) in $(DIR)..."
	@./bin/migrate-v2 -config $(DIR) -from $(FROM) -to $(TO)

# Dry run migration
# Usage: make dry-run DIR=./test-configs
dry-run: build
	@echo "Dry run migration in $(DIR)..."
	@./bin/migrate-v2 -config $(DIR) -dry-run -preview

# Help
help:
	@echo "Available targets:"
	@echo "  build           - Build the migrate-v2 binary"
	@echo "  test            - Run all tests (unit + integration)"
	@echo "  test-unit       - Run unit tests only"
	@echo "  test-integration - Run all integration tests"
	@echo "  test-resource   - Run integration tests for a specific resource (RESOURCE=name)"
	@echo "  test-case       - Run specific test case (RESOURCE=name CASE=basic|complete)"
	@echo "  test-verbose    - Run all tests with verbose output"
	@echo "  test-coverage   - Generate test coverage report"
	@echo "  benchmark       - Run benchmark tests"
	@echo "  clean           - Remove build and test artifacts"
	@echo "  validate-fixtures - Validate all test fixtures are valid HCL"
	@echo "  update-golden   - Update expected v5 outputs (USE WITH CAUTION)"
	@echo "  migrate         - Run migration (FROM=v4 TO=v5 DIR=path)"
	@echo "  dry-run         - Dry run migration (DIR=path)"
	@echo "  help            - Show this help message"