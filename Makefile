TARGET := folders
BINDIR := bin
SRCDIR := folder
MAINFILE := main.go
SRCFILES := $(wildcard folder/*.go)

.PHONY: all build build_run test clean unit_tests e2e_tests

all: build test

build: ./$(BINDIR)/$(TARGET)

run:
	@go run main.go

build_run : build test
	./$(BINDIR)/$(TARGET)

test: unit_tests

unit_tests: $(MAINFILE) $(SRCFILES)
	@go test ./folder/get_folder_test.go
	@go test ./folder/move_folder_test.go

e2e_tests: $(BINDIR)/$(TARGET)
	@echo E2E Tests in progress

./$(BINDIR)/$(TARGET): $(MAINFILE) $(SRCFILES)
	@mkdir -p $(BINDIR)
	go build -o $(BINDIR)/$(TARGET)

clean:
	@rm -rf bin
