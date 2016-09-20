BUILD_DIR = build
MARGO_SRC = main.go
MARGO_CLI_SRC = tool/margo-api-cli/main.go

all: clean get-deps margo margo-cli test

get-deps:
	@go get -t ./...

margo: $(MARGO_SRC)
	@go build -o $(BUILD_DIR)/margo

margo-cli: $(MARGO_CLI_SRC)
	@go build -o $(BUILD_DIR)/margo-cli $(MARGO_CLI_SRC)

clean:
	@rm -rf $(BUILD_DIR)

test:
	@go test -v -cover ./...
