package transforms

import (
	"github.com/headroom-ai/headroom-go/pkg/models"
	"github.com/headroom-ai/headroom-go/pkg/tokenizers"
)

// BasicCompressor is a simple compressor that truncates content for MVP demonstration.
type BasicCompressor struct{}

func (b *BasicCompressor) Name() string { return "basic_compressor" }

func (b *BasicCompressor) Apply(messages []models.Message, tokenizer tokenizers.Tokenizer, config models.CompressConfig) (models.TransformResult, error) {
	newMessages := make([]models.Message, len(messages))
	transformsApplied := []string{}
	
	tokensBefore := tokenizer.CountMessages(messages)

	// Simple logic: if a message is over MinTokensToCompress, truncate it by half.
	for i, m := range messages {
		msgTokens := tokenizer.CountText(m.Content)
		if msgTokens > config.MinTokensToCompress {
			// Truncate content
			halfLen := len(m.Content) / 2
			newMessages[i] = models.Message{
				Role:    m.Role,
				Content: m.Content[:halfLen] + "... [headroom truncated]",
			}
			transformsApplied = append(transformsApplied, "basic_truncate")
		} else {
			newMessages[i] = m
		}
	}

	tokensAfter := tokenizer.CountMessages(newMessages)

	return models.TransformResult{
		Messages:          newMessages,
		TokensBefore:      tokensBefore,
		TokensAfter:       tokensAfter,
		TransformsApplied: transformsApplied,
		Timing:            make(map[string]float64),
	}, nil
}

func (b *BasicCompressor) ShouldApply(messages []models.Message, tokenizer tokenizers.Tokenizer, config models.CompressConfig) bool {
	return true
}

// ContentRouter routes messages to specialized compressors.
type ContentRouter struct {
	compressor Transform
}

func NewContentRouter(compressor Transform) *ContentRouter {
	return &ContentRouter{compressor: compressor}
}

func (c *ContentRouter) Name() string {
	return "content_router"
}

func (c *ContentRouter) Apply(messages []models.Message, tokenizer tokenizers.Tokenizer, config models.CompressConfig) (models.TransformResult, error) {
	// For MVP, we route everything to the provided compressor.
	// In the real version, this would split the message list by content type.
	return c.compressor.Apply(messages, tokenizer, config)
}

func (c *ContentRouter) ShouldApply(messages []models.Message, tokenizer tokenizers.Tokenizer, config models.CompressConfig) bool {
	return true
}
