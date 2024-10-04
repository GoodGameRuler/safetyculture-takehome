TARGET := bin/program

.PHONY: all test clean unit_tests e2e_tests

run: test
	@go run main.go

build_run : $(TARGET) test
	./bin/program

test: unit_tests e2e_tests

unit_tests:
	@go test ./folder/get_folder_test.go
	@go test ./folder/move_folder_test.go

e2e_tests: $(TARGET)
	@echo in progress

$(TARGET):
	@mkdir -p bin
	@go build main.go bin/program

clean:
	@rm -rf bin

_run:
	@echo run in progress
