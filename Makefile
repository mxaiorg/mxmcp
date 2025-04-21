linux-intel:
	env GOOS=linux GOARCH=amd64 go build -v -o mxmcp *.go

dev-mac:
	go build -v \
	-ldflags="-X main.ApiHost=http://localhost:9080" \
	-o mxmcp-dev *.go

mac-arm:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
	go build -v \
	-ldflags="-X main.ApiHost=https://lab4-api.mxhero.com" \
	-o bin/mxmcp-mac-arm *.go

mac-intel:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
	go build -v \
	-ldflags="-X main.ApiHost=https://lab4-api.mxhero.com" \
	-o bin/mxmcp-mac-intel *.go

windows-intel:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
	go build -v \
	-ldflags="-X main.ApiHost=https://lab4-api.mxhero.com" \
	-o bin/mxmcp.exe *.go
