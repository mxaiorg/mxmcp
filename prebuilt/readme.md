# Prebuilt Binaries

This folder contains pre-built binaries to facilitate deployment.

## Mac ARM (Apple Silicon—M1, M2 ...) 

**mxmcp-mac-arm-installer.pkg**

* Double-click the package to install.
* md5sum
  * 36f344c51eb093ed6d87c4d83bf031e2 mxmcp-mac-arm-installer.pkg

## Mac AMD (Apple Intel)

**mxmcp-mac-intel-installer.pkg**

* Double-click the package to install.
* md5sum
  * 50ec56610a4f3f932100542a7ff125fe mxmcp-mac-intel-installer.pkg

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