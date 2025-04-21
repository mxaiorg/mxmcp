# MCP Server for mxHERO Multi-Account Email Search

## Description

This MCP (model context protocol) server is a Go project that provides access to mxHERO's multi-account email search service.

The Model Context Protocol (MCP) is a framework designed to standardize the way models interact with various data sources and services. In this project, MCP is used to facilitate seamless integration to emails captured by mxHERO Mail2Cloud. Mail2Cloud is designed to selectively capture emails from one or more accounts. The selection of emails can be finely controlled by powerful filters examining almost any aspect of messages and their attachments.

For more information about mxHERO's multi-email account service, including architecture, optimizations, etc. [read here](https://mxhero.com).

**Why Go for MCP deployment**

Unlike Python or Javascript MCPs, Go compiles to native binary. Once compiled for a target architecture (e.g., Mac ARM, Windows Intel) and installed, no additional dependencies are required on the user's device.


### Alternate versions

A Python version is in development and will be linked here.

## Tools implemented

## `email_search`
Search stored emails

**Parameters**
- `query` (str): Email search query

**Returns** JSON of search results


## Requirements

- GO 1.22 or higher ([download](https://go.dev/doc/install))
- mxHERO Vector Search credentials (token)
  - A demo token can be obtained at https://lab4-api.mxhero.com/demo_signup
  - For production tokens, contact mxHERO.

## Installation
1. Clone the repository
```sh
git clone https://github.com/mxaiorg/mxmcp
cd mxmcp
go mod tidy
```

2. Compile

Be sure to compile for the architecture of the user's computer. You will need to match the operating system and processor architecture. The included Makefile provides for a few of the most common.

| OS      | Architecture              | Make command       |
|---------|---------------------------|--------------------|
| Windows | Intel                     | make windows-intel |
| Mac     | Arm (Mac silicon - M1...) | make mac-arm       |
| Mac     | Intel                     | make mac-intel     |
| Linux   | Intel                     | make linux-intel   |

For more operating systems and architectures see Go compilation documentation.

**Example build**

```shell
make mac-arm
```

After `make` is run it will place the program (binary) in the `bin`folder. Copy this binary to the user's computer and see the configuration instructions below.

**Note**
* Some platforms, like MacOS, will require additional permissions before allowing the program to be run on another machine.

**Prebuilt Binaries**
* For convenience the `prebuilt` contains prebuilt binaries and signed installation packages. 

## Installation

If not installing with an installation package do the following:

1. Copy the binary (of matching operating systems and architecture) to the user's computer. Place the file somewhere the user has permissions to access. For example, the user's home directory.


2. Ensure the user has permission to run the program (execute)
   - For example, on Mac & Linux `chmod 755 mxmcp`


### Configuring for Claude Desktop

The following is an example configuration JSON for common clients (e.g., Claude). For details of installing MCPs in Claude see https://modelcontextprotocol.io/quickstart/user

1. Edit your `claude_desktop_config`.json
   * You make need to create the file if it does not already exist.


2. Add the following JSON below, where:
   * Note mxhero-mcp-server JSON should be added alongside any other MCP servers of your configuration.
   * Be sure to put the full path as the command value. For example:
     * `/Users/bob/mxmcp`
   * Parameters are:
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