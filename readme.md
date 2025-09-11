# MCP Server for mxHERO Multi-Account Email Search

## Description

This MCP (model context protocol) server is a Go project that provides access to mxHERO's multi-account email search service.

The Model Context Protocol (MCP) is a framework designed to standardize the way models interact with various data sources and services. In this project, MCP is used to facilitate seamless integration to mxHERO Mail2Cloud Advanced. Mail2Cloud Advanced is a high performance data service for a company's email data. Mail2Cloud Advanced connects to company email services and optimizes the content for fast, scalable and secure access by AI solutions.

### Architecture
Mail2Cloud is designed to selectively capture emails from one or more accounts. The selection of emails can be finely controlled by powerful filters examining any aspect of messages and their attachments. Captured emails are then optimized and stored into an isolated tenant in a vector database designed for email related searches. This MCP accesses the stored emails in the tenant through authenticated access credentials.

### Advantages
Solutions built with Mail2Cloud Advanced MCP outperforms other AI solutions with regards to email data search & knowledge recovery ([study](https://medium.com/datadriveninvestor/ai-email-retrieval-benchmark-how-purpose-built-ai-tools-outperform-generic-solutions-6fcd6d560c8f))
* Provides secure links to original emails (safe from accidental user deletion, etc.)
* Allows LLMs to search massive email repositories, far beyond their context window restraints.

### Demo Accounts

To facilitate exploration of this MCP, mxHERO provides demo accounts that are pre-loaded with thousands of emails. More about the demo emails can be found [here](https://mxhero.helpjuice.com/en_US/mxhero-ai/demo-account-for-ai-testing).

For more information see: [mxHERO Mail2Cloud Advanced](https://www.mxhero.com/advanced-ai) multi-email account service, including [architecture](https://mxhero.helpjuice.com/en_US/mxhero-ai/mxmcp#architecture-8), and optimizations.

### Why Go for MCP deployment

Unlike Python or Javascript MCPs, Go compiles to native static binary. Once compiled for a target architecture (e.g., Mac ARM, Windows Intel) and installed, no additional dependencies (software) are required on the user's device.

## Quickstart

1. To get started quickly, see the prebuilt binaries section below for your operating system and machine architecture.
2. Follow the installation instructions.

## Alternate versions

A Python version can be found [here](https://github.com/mxaiorg/mxmcp-py).

### Streamable HTTP
This MCP repo is the 'stdio' variant. HTTP options exist at the following addresses:
* [https://lab4-api.mxhero.com/mcp/connect](https://lab4-api.mxhero.com/mcp/connect)
  (streamable HTTP)
* [https://lab4-api.mxhero.com/mcp/sse](https://lab4-api.mxhero.com/mcp/sse) (Legacy SSE)

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

For example: `cp bin/mxmcp-mac-arm ~user`

**Note**
* Some platforms, like MacOS, will require additional permissions before allowing the program to be run on another machine.

## Prebuilt Binaries
For convenience the `prebuilt` folder contains prebuilt binaries and signed installation packages. See the "readme" file in that folder for more information.


## Installation

If **NOT** installing with an installation package, do the following:

1. Copy the binary (of matching operating systems and architecture) to the user's computer. Place the file somewhere the user has permission to access. For example, the user's home directory.


2. Ensure the user has permission to run the program (execute)
   - For example, on Mac & Linux `chmod 755 mxmcp`


### Configuring for Claude Desktop

The following is an example configuration JSON for common clients (e.g., Claude). For details of installing MCPs in Claude see https://modelcontextprotocol.io/quickstart/user

1. Edit your `claude_desktop_config`.json
   * You may need to create the file if it does not already exist.


2. Add the following JSON below, where:
   * Note mxhero-mcp-server JSON should be added alongside any other MCP servers of your configuration.
   * Be sure to put the **full path** as the command value. For example:
     * `/Users/bob/mxmcp-mac-arm`
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

3. If Claude is running, restart it.

## Notice

Using this client against the hosted service requires an account/API key and is governed by our ToS.