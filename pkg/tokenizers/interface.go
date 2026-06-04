package tokenizers

import "github.com/headroom-ai/headroom-go/pkg/models"

// Tokenizer defines the interface for token counting.
type Tokenizer interface {
	CountMessages(messages []models.Message) int
	CountText(text string) int
	Model() string
}
