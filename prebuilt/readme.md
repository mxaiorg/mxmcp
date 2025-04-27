# Prebuilt Binaries

This folder contains pre-built binaries to facilitate deployment.

## Mac ARM (Apple Silicon—M1, M2 ...) 

**mxmcp-mac-arm-installer.pkg**

* Double-click the package to install.
* md5sum
  * 48884025beb57e6d8d6a693818637d36 mxmcp-mac-arm-installer.pkg

## Mac AMD (Apple Intel)

**mxmcp-mac-intel-installer.pkg**

* Double-click the package to install.
* md5sum
  * 8ff9e68af34ff5f8f3796ba5c68ff2e4 mxmcp-mac-intel-installer.pkg

## Configuration

### Mac

The Mac package installs mxmcp-mac-arm/intel to your `/usr/local/bin` directory. As such, set the command field in your configuration to:

* Mac ARM: `/usr/local/bin/mxmcp-mac-arm`
* Mac Intel: `/usr/local/bin/mxmcp-mac-intel`

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