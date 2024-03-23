# Makefile for running Fiber project

# Variables
APP_NAME := go-fiber-boilerplate
BUILD_DIR := bin
SRC_FILES := $(shell find . -name "*.go" -type f)
MAIN_PACKAGE := main.go

# Build the project
build: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME): $(SRC_FILES)
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PACKAGE)

# Run the project
run: build
	$(BUILD_DIR)/$(APP_NAME)

# Clean the project
clean:
	@rm -rf $(BUILD_DIR)

.PHONY: build run clean
