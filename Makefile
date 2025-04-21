linux-intel:
	env GOOS=linux GOARCH=amd64 go build -v -o mxmcp *.go

dev-mac:
	go build -v \
	-ldflags="-X main.ApiHost=http://localhost:9080" \
	-o demo-mxmcp-dev *.go

mac-arm:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
	go build -v \
	-ldflags="-X main.ApiHost=https://lab4-api.mxhero.com" \
	-o demo-mxmcp-dev *.go

mac-intel:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
	go build -v \
	-ldflags="-X main.ApiHost=https://lab4-api.mxhero.com" \
	-o bin/demo-mxmcp-intel *.go

windows-intel:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
	go build -v \
	-ldflags="-X main.ApiHost=https://lab4-api.mxhero.com" \
	-o bin/demo-mxmcp.exe *.go
