# Headroom Go

A high-performance Golang re-implementation of the [Headroom](https://github.com/headroom-ai/headroom) token reduction proxy and SDK.

Headroom helps you reduce the cost and latency of LLM applications by intelligently compressing message history before sending it to the model. It uses specialized compressors for different content types (JSON, Code, Text) to shrink context windows without losing meaning.

## Features (Current MVP)

- **Transform Pipeline:** Sequential orchestration of message transformations.
- **Cache Aligner:** Stabilizes message prefixes to maximize provider KV cache hits.
- **Content Router:** Dispatches content chunks to specialized compressors.
- **Basic Compressor:** Truncates long messages while preserving context.
- **MCP Server:** Built-in Model Context Protocol server to expose compression as an AI tool.

## Project Structure

```text
.
├── cmd/
│   └── mcp/            # MCP Server entry point
├── pkg/
│   ├── compress/       # High-level Public API
│   ├── models/         # Shared data structures
│   ├── tokenizers/     # Token counting interfaces & logic
│   └── transforms/     # Pipeline and Transform implementations
└── GEMINI.md           # AI Architectural Guardrails
```

## Getting Started

### Prerequisites

- Go 1.22+

### Installation

```bash
go get github.com/headroom-ai/headroom-go
```

## Agent Integration (MCP)

Headroom Go includes a built-in MCP server to allow AI agents to compress their own context.

### Build the MCP Server

```bash
make build
```

### Configure for Claude Desktop / Gemini CLI

Add to your `mcp_config.json` (using the absolute path to the binary):

```json
{
  "mcpServers": {
    "headroom-go": {
      "command": "{path-to-headroom-mcp-cli}",
      "args": []
    }
  }
}
```

## Roadmap

- [x] MVP Pipeline & Basic Truncation
- [x] MCP Server Integration
- [ ] **SmartCrusher:** JSON array deduplication
- [ ] **CodeCompressor:** AST-aware code shrinking
- [ ] **Kompress:** ML-based text compression (ONNX)
- [ ] **Proxy Server:** Standalone HTTP proxy for Anthropic/OpenAI

## License

Apache-2.0
