package tokenizers

import (
	"github.com/headroom-ai/headroom-go/pkg/models"
	"unicode/utf8"
)

// SimpleTokenizer is a basic character-based tokenizer for MVP testing.
type SimpleTokenizer struct {
	model string
}

func NewSimpleTokenizer(model string) *SimpleTokenizer {
	return &SimpleTokenizer{model: model}
}

func (t *SimpleTokenizer) CountMessages(messages []models.Message) int {
	total := 0
	for _, m := range messages {
		total += t.CountText(m.Role)
		total += t.CountText(m.Content)
		total += 3 // Approximate overhead per message
	}
	return total
}

func (t *SimpleTokenizer) CountText(text string) int {
	// Crude estimation: ~4 chars per token for English
	chars := utf8.RuneCountInString(text)
	if chars == 0 {
		return 0
	}
	tokens := chars / 4
	if tokens == 0 {
		return 1
	}
	return tokens
}

func (t *SimpleTokenizer) Model() string {
	return t.model
}
