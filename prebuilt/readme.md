# Prebuilt Binaries

This folder contains pre-built binaries to facilitate deployment.

## Mac ARM (Apple Silicon—M1, M2 ...) 

**mxmcp-mac-arm-installer.pkg**

* Double-click the package to install.
* md5sum
  * f051795c2a24b300eb91a4fd375dded7  prebuilt/mxmcp-mac-arm-installer.pkg

## Mac AMD (Apple Intel)

**mxmcp-mac-intel-installer.pkg**

* Double-click the package to install.
* md5sum
  * 1c32c6393dc4cb72a25a734981e166f6  prebuilt/mxmcp-mac-intel-installer.pkg

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