TARGET := folders
BINDIR := bin
SRCDIR := folder
MAINFILE := main.go

.PHONY: all build build_run test clean unit_tests e2e_tests

build: test $(BINDIR)/$(TARGET)

run: test
	@go run main.go

build_run : build test
	./$(BINDIR)/$(TARGET)

test: unit_tests e2e_tests

unit_tests:
	@go test ./folder/get_folder_test.go
	@go test ./folder/move_folder_test.go

e2e_tests: $(BINDIR)/$(TARGET)
	@echo E2E Tests in progress

$(BINDIR)/$(TARGET): $(MAINFILE) $(wildcard $(SRCDIR)/**/*.go)
	@mkdir -p $(BINDIR)
	go build -o $(BINDIR)/$(TARGET)

clean:
	@rm -rf bin
