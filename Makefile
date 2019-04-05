BUILD_CONTEXT := ./build

.PHONY: go-test
go-test:
	@echo "Run all project tests..."
	go test -p 1 ./...

.PHONY: bin
bin: clean
	@echo "Build project binaries..."
	GOOS=linux GOARCH=386 go build -v -o $(BUILD_CONTEXT)/lucas_linux_386
	GOOS=darwin GOARCH=386 go build -v -o $(BUILD_CONTEXT)/lucas_darwin_386
	GOOS=linux GOARCH=386 go build -v -o $(BUILD_CONTEXT)/lucas_windows_386

.PHONY: clean
clean:
	rm -rf $(BUILD_CONTEXT)/

.PHONY: test
test: go-test

.PHONY: run-cern
run-cern:
	go run lucas.go http://info.cern.ch
