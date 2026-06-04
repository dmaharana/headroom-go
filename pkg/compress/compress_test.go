package compress

import (
	"testing"

	"github.com/headroom-ai/headroom-go/pkg/models"
	"github.com/headroom-ai/headroom-go/pkg/transforms"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompress(t *testing.T) {
	messages := []models.Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: "This is a very long message that should be truncated by our basic compressor because it exceeds the minimum token count threshold. " +
			"Repeat: This is a very long message that should be truncated by our basic compressor because it exceeds the minimum token count threshold. " +
			"Repeat: This is a very long message that should be truncated by our basic compressor because it exceeds the minimum token count threshold."},
	}

	pipeline := transforms.NewTransformPipeline(
		transforms.NewCacheAligner(),
		transforms.NewContentRouter(&transforms.BasicCompressor{}),
	)
	compressor := NewCompressor(pipeline)

	config := models.NewDefaultCompressConfig()
	config.MinTokensToCompress = 20 // Lower threshold for test

	result, err := compressor.Compress(messages, "gpt-4", &config)
	require.NoError(t, err)

	assert.Equal(t, len(messages), len(result.Messages))
	assert.Greater(t, result.TokensBefore, result.TokensAfter)
	assert.Greater(t, result.TokensSaved, 0)
	assert.Greater(t, result.CompressionRatio, 0.0)
	assert.Contains(t, result.TransformsApplied, "basic_truncate")
}
