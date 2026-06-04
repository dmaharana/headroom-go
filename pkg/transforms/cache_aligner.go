package transforms

import (
	"github.com/headroom-ai/headroom-go/pkg/models"
	"github.com/headroom-ai/headroom-go/pkg/tokenizers"
)

// CacheAligner ensures the prefix of the message list remains stable to improve provider cache hits.
type CacheAligner struct {
}

func NewCacheAligner() *CacheAligner {
	return &CacheAligner{}
}

func (c *CacheAligner) Name() string {
	return "cache_aligner"
}

func (c *CacheAligner) Apply(messages []models.Message, tokenizer tokenizers.Tokenizer, config models.CompressConfig) (models.TransformResult, error) {
	// In Python, this transform might re-order or stabilize messages.
	// For MVP, we'll just acknowledge it ran and return the messages as-is.
	// Real implementation would ensure 'protect_recent' and other flags are respected.
	
	tokens := tokenizer.CountMessages(messages)
	
	return models.TransformResult{
		Messages:          messages,
		TokensBefore:      tokens,
		TokensAfter:       tokens,
		TransformsApplied: []string{"cache_aligner"},
		Timing:            make(map[string]float64),
	}, nil
}

func (c *CacheAligner) ShouldApply(messages []models.Message, tokenizer tokenizers.Tokenizer, config models.CompressConfig) bool {
	return true
}
