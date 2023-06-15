BINARY_NAME=ezweeb

build:
	@go build -o ./bin/${BINARY_NAME} ./cmd/ezweeb/ezweeb.go

build-dist:
ifeq ($(OS),Windows_NT)
	$(info Building for Windows)
	@go build -o bin/${BINARY_NAME}-windows ./cmd/ezweeb/ezweeb.go
	$(info Done!)
	$(info ~~~)
	$(info Building for Linux)
	@set GOARCH=amd64 
	@set GOOS=linux
	@go build -o bin/${BINARY_NAME}-linux ./cmd/ezweeb/ezweeb.go
	$(info Done!)
	$(info ~~~)
	$(info Building for MacOS-AMD)
	@set GOARCH=amd64 
	@set GOOS=darwin
	@go build -o bin/${BINARY_NAME}-darwin_amd ./cmd/ezweeb/ezweeb.go
	$(info Done!)
	$(info ~~~)
	$(info Building for MacOS-ARM)
	@set GOARCH=arm64 
	@set GOOS=darwin
	@go build -o bin/${BINARY_NAME}-darwin_arm ./cmd/ezweeb/ezweeb.go
	$(info Done!)
else
	$(info Building for Linux)
	@GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux ./cmd/ezweeb/ezweeb.go
	$(info Done!)
	$(info ~~~)
	$(info Building for MacOS-AMD)
	@GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin_amd ./cmd/ezweeb/ezweeb.go
	$(info Done!)
	$(info ~~~)
	$(info Building for MacOS-ARM)
	@GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin_arm ./cmd/ezweeb/ezweeb.go
	$(info Done!)
	$(info ~~~)
	$(info Building for Windows)
	@GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows ./cmd/ezweeb/ezweeb.go
	$(info Done!)
endif


run: build
	@./bin/${BINARY_NAME}
	

clean:
	@go clean
ifeq ($(OS),Windows_NT)
	@powershell "Remove-Item .\bin\*" 
else
	@rm -f ./bin/*
endif