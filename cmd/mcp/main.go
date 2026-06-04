package main

import (
	"context"
	"fmt"
	"os"

	"github.com/headroom-ai/headroom-go/pkg/compress"
	"github.com/headroom-ai/headroom-go/pkg/models"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Headroom Go",
		"1.0.0",
	)

	// Add the compress_messages tool
	s.AddTool(mcp.NewTool("compress_messages",
		mcp.WithDescription("Compress a list of LLM messages to reduce token count while preserving meaning."),
		mcp.WithArray("messages",
			mcp.Required(),
			mcp.Description("The list of messages to compress."),
		),
		mcp.WithString("model", mcp.Description("The model ID used for token counting (e.g., gpt-4o, claude-3-5-sonnet).")),
	), compressHandler)

	// Start the server on stdio
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func compressHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Type assert Arguments to map[string]interface{}
	args, ok := request.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("invalid arguments format"), nil
	}

	// Parse messages
	messagesArg, ok := args["messages"]
	if !ok {
		return mcp.NewToolResultError("missing messages argument"), nil
	}

	modelArg, ok := args["model"]
	model := "gpt-4o"
	if ok {
		model, _ = modelArg.(string)
	}

	// Convert raw arguments to models.Message
	rawMessages, ok := messagesArg.([]interface{})
	if !ok {
		return mcp.NewToolResultError("messages must be an array"), nil
	}

	messages := make([]models.Message, len(rawMessages))
	for i, raw := range rawMessages {
		m, ok := raw.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError(fmt.Sprintf("message %d must be an object", i)), nil
		}
		
		role, _ := m["role"].(string)
		content, _ := m["content"].(string)
		
		messages[i] = models.Message{
			Role:    role,
			Content: content,
		}
	}

	// Perform compression using the DefaultCompress helper
	result, err := compress.DefaultCompress(messages, model)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("compression failed: %v", err)), nil
	}

	// Format response
	return mcp.NewToolResultText(fmt.Sprintf(
		"Compression Complete:\n- Tokens Before: %d\n- Tokens After: %d\n- Tokens Saved: %d\n- Ratio: %.2f%%\n\nTransforms Applied: %v",
		result.TokensBefore,
		result.TokensAfter,
		result.TokensSaved,
		result.CompressionRatio*100,
		result.TransformsApplied,
	)), nil
}
