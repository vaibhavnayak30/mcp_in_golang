# Math MCP Server

A standalone Model Context Protocol (MCP) server implemented in Go that provides basic mathematical operations.

## Features

* 🔌 **MCP Protocol Support**: Implements the Model Context Protocol for seamless integration with MCP clients.

* 🌐 **Dual Endpoints**: Supports both HTTP JSON-RPC and Server-Sent Events (SSE).

* ➕➖✖️➗ **Mathematical Operations**: Provides 5 basic math operations as tools.

* 🚀 **REST API**: Additional HTTP endpoints for health checks and configuration.

* 📚 **Swagger Documentation**: Auto-generated API documentation.

## Available Tools

| **Tool** | **Description** | **Parameters** |
| :--------- | :------------------------------------------- | :-------------------------- |
| `add`      | Add two numbers together                     | `a` (float64), `b` (float64) |
| `subtract` | Subtract second number from first number     | `a` (float64), `b` (float64) |
| `multiply` | Multiply two numbers together                | `a` (float64), `b` (float64) |
| `divide`   | Divide first number by second number         | `a` (float64), `b` (float64) |
| `power`    | Calculate base raised to the power of exponent | `base` (float64), `exponent` (float64) |

## Installation

### Prerequisites

* Go 1.24 or later

### Build from Source

1.  Clone the repository:
    ```bash
    git clone [https://github.com/vaibhavnayak30/mcp_in_golang.git](https://github.com/vaibhavnayak30/mcp_in_golang.git)
    ```
2.  Navigate to the project directory:
    ```bash
    cd math-mcp
    ```
3.  Build the project:
    ```bash
    make build
    ```

### Run the Server

To run the server, use the `make run` command or execute the binary directly:

```bash
# Using make
make run

# Or directly
./bin/math-mcp