# mxMCP

This MCP provides access to mxHERO's Advanced vector email search service.

## Configuration file

The following is an example configuration JSON for common clients (e.g., Claude).
* mxhero-mcp-server should be added alongside any other MCP servers of your configuration. 
* '-t' (token) parameter is required.
* '-d' is an optional custom tool description.

```json
{
  "mcpServers": {
    "mxhero-mcp-server": {
      "command": "<full path to mxmcp>",
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