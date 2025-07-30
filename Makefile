.PHONY: test test-unit test-integration

test: test-unit test-integration

test-unit:
		@echo "Running unit tests..."
		@go test -v ./internal/tests/unit/...

test-integration:
		@echo "Running integration tests..."
		@go test -v ./internal/tests/integration/...

test-coverage:
		@echo "Running tests with coverage..."
		@go test -coverprofile=coverage.out ./...
		@go tool cover -html/coverage.out