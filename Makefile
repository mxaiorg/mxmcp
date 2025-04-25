linux-intel:
	env GOOS=linux GOARCH=amd64 go build -v -o mxmcp *.go

mac-arm-local:
	go build -v \
	-ldflags="-X main.ApiHost=http://localhost:9080" \
	-o bin/mxmcp-local *.go

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

# Local mac testing
mac-local-install: mac-arm-local
	chmod +x bin/mxmcp-local
	cp bin/mxmcp-local /usr/local/bin/mxmcp-mac-arm

mac-arm-install: mac-arm
	chmod +x bin/mxmcp-mac-arm
	cp bin/mxmcp-mac-arm /usr/local/bin/mxmcp-mac-arm
