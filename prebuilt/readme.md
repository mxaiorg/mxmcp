# Prebuilt Binaries

This folder contains pre-built binaries to facilitate deployment.

### Mac ARM (Apple Silicon—M1, M2 ...) 

mxmcp-mac-arm-installer.pkg

* Double-click the package to install.
* md5sum
  * 17c66a6921ccc63f1df2356aa677800c  mxmcp-mac-arm-installer.pkg

#### Configuration

The Mac package installs mxmcp-mac-arm to your `/usr/local/bin` directory. As such, set the command field in your configuration to:

`/usr/local/bin/mxmcp-mac-arm`

For example:

```json
{
  "mcpServers": {
    "mxhero-mcp-server": {
      "command": "/usr/local/bin/mxmcp-mac-arm",
      "args": [
        "-t",
        "<token>",
        "-d",
        "<optional custom tool description>"
      ]
    }
  }
}
```