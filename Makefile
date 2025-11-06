.PHONY: fmt
fmt: 
	@gofumpt -l -w .

build: test
	@echo "building binaries...."
	@GOOS=linux GOARCH=amd64 go build -o bin/wsc_lx_amd64 cmd/checker/main.go
	@GOOS=linux GOARCH=arm64 go build -o bin/wsc_lx_arm64 cmd/checker/main.go
	@GOOS=darwin GOARCH=amd64 go build -o bin/wsc_dw_amd64 cmd/checker/main.go
	@GOOS=darwin GOARCH=arm64 go build -o bin/wsc_dw_arm64 cmd/checker/main.go
	@GOOS=windows GOARCH=amd64 go build -o bin/wsc_wi_amd64 cmd/checker/main.go
	@GOOS=windows GOARCH=arm64 go build -o bin/wsc_wi_arm64 cmd/checker/main.go
	@echo "done."

test:
	go test -v ./...