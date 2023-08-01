BINARY_NAME=ezweeb

build:
	@go mod tidy 
ifeq ($(OS),Windows_NT)
	@go build -o ./bin/${BINARY_NAME}.exe ./cmd/ezweeb/ezweeb.go
else
	@go build -o ./bin/${BINARY_NAME} ./cmd/ezweeb/ezweeb.go
endif

build-dist:
	@go mod tidy 
ifeq ($(OS),Windows_NT)
	@powershell "echo \"building for windows\"; go build -o bin/${BINARY_NAME}_windows.exe ./cmd/ezweeb/ezweeb.go; echo done"
	@powershell "echo \"building for linux-amd\"; go env -w GOOS=linux; go build -o bin/${BINARY_NAME}_linux-amd ./cmd/ezweeb/ezweeb.go; echo done"
	@powershell "echo \"building for darwin-amd\"; go env -w GOOS=darwin; go build -o bin/${BINARY_NAME}_darwin-amd ./cmd/ezweeb/ezweeb.go; echo done"
	@powershell "echo \"building for linux-arm\"; go env -w GOOS=linux; go env -w GOARCH=arm64; go build -o bin/${BINARY_NAME}_linux-arm ./cmd/ezweeb/ezweeb.go; echo done"
	@powershell "echo \"building for darwin-arm\"; go env -w GOOS=darwin; go env -w GOARCH=arm64; go build -o bin/${BINARY_NAME}_darwin-arm ./cmd/ezweeb/ezweeb.go; echo done"
	@powershell "go env -w GOOS=windows; go env -w GOARCH=amd64"
else
	@echo "building for linux-amd";GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}_linux-x64 ./cmd/ezweeb/ezweeb.go; echo "done"
	@echo "building for linux-arm";GOARCH=arm64 GOOS=linux go build -o bin/${BINARY_NAME}_linux-arm ./cmd/ezweeb/ezweeb.go; echo "done"
	@echo "building for darwin-amd";GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}_darwin-x64 ./cmd/ezweeb/ezweeb.go; echo "done"
	@echo "building for darwin-arm";GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}_darwin-arm ./cmd/ezweeb/ezweeb.go; echo "done"
	@echo "building for windows";GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows.exe ./cmd/ezweeb/ezweeb.go; echo "done"
endif


run: build
	@go mod tidy 
	@./bin/${BINARY_NAME}
	

clean:
	@go clean
ifeq ($(OS),Windows_NT)
	@powershell "Remove-Item .\bin\*" 
else
	@rm -f ./bin/*
endif