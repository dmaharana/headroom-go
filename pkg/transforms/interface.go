package transforms

import (
	"github.com/headroom-ai/headroom-go/pkg/models"
	"github.com/headroom-ai/headroom-go/pkg/tokenizers"
)

// Transform defines the interface for a message transformation step.
type Transform interface {
	Name() string
	Apply(messages []models.Message, tokenizer tokenizers.Tokenizer, config models.CompressConfig) (models.TransformResult, error)
	ShouldApply(messages []models.Message, tokenizer tokenizers.Tokenizer, config models.CompressConfig) bool
}
