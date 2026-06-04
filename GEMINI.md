# Headroom Go - Project Instructions

You are working in the `headroom-go` repository, a Golang re-implementation of the Python Headroom token reduction proxy and SDK.

## Architecture & Layout
- **Pattern:** Standard Go Layout.
- **Entrypoints (`cmd/`):** All executable binaries live here (e.g., `cmd/mcp`, `cmd/proxy`). They should contain minimal logic, primarily handling dependency injection, flag parsing, and starting the server/service.
- **MCP Server:** The MCP server (`cmd/mcp`) is designed to be run as a compiled binary (typically built to `bin/headroom-mcp`) rather than via `go run` for production/integration usage.
- **Internal Logic (`internal/`):** Use for code that should NOT be imported by external users.
- **Public SDK (`pkg/`):** Core library code. Specifically:
  - `pkg/models`: Shared domain structures (`Message`, `CompressConfig`, `TransformResult`).
  - `pkg/transforms`: The core `TransformPipeline` and individual `Transform` implementations (e.g., `CacheAligner`, `ContentRouter`).
  - `pkg/tokenizers`: Token counting interfaces and implementations.
  - `pkg/compress`: The high-level public API.

## Core Concepts
- **Transforms:** Every operation that modifies a message list is a `Transform`. They MUST implement the `transforms.Transform` interface (`Apply()` and `ShouldApply()`).
- **Pipeline:** Transforms are executed sequentially by the `TransformPipeline`.
- **Immutability (Soft):** Transforms should avoid mutating the incoming message array directly; they should allocate new arrays or structures when modifying content, tracking `TokensBefore` and `TokensAfter`.

## Coding Conventions
- **Dependency Injection:** We use manual constructor injection (e.g., `NewTransformPipeline(...)`) for the library core, and plan to use `google/wire` for large application binaries in `cmd/`.
- **Testing:** Use `github.com/stretchr/testify/assert` and `require` for all tests. Co-locate `_test.go` files with their packages.
- **Error Handling:** Return explicit errors. Do not panic in library code (`pkg/`).

## Agent Role
When modifying this codebase, you must respect the pipeline architecture. If you add a new compression strategy (e.g., Markdown compression), you MUST implement it as a new struct satisfying `transforms.Transform` and register it within the appropriate router or pipeline.
