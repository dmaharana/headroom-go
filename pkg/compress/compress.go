package compress

import (
	"github.com/headroom-ai/headroom-go/pkg/models"
	"github.com/headroom-ai/headroom-go/pkg/tokenizers"
	"github.com/headroom-ai/headroom-go/pkg/transforms"
)

// Compressor is the main service for compressing messages.
type Compressor struct {
	pipeline *transforms.TransformPipeline
}

func NewCompressor(pipeline *transforms.TransformPipeline) *Compressor {
	return &Compressor{
		pipeline: pipeline,
	}
}

// Compress takes a list of messages and returns a compressed version.
func (c *Compressor) Compress(messages []models.Message, model string, config *models.CompressConfig) (models.CompressResult, error) {
	if config == nil {
		d := models.NewDefaultCompressConfig()
		config = &d
	}

	tokenizer := tokenizers.NewSimpleTokenizer(model)
	
	result, err := c.pipeline.Apply(messages, tokenizer, *config)
	if err != nil {
		return models.CompressResult{
			Messages: messages,
		}, err
	}

	tokensSaved := result.TokensBefore - result.TokensAfter
	ratio := 0.0
	if result.TokensBefore > 0 {
		ratio = float64(tokensSaved) / float64(result.TokensBefore)
	}

	return models.CompressResult{
		Messages:          result.Messages,
		TokensBefore:      result.TokensBefore,
		TokensAfter:       result.TokensAfter,
		TokensSaved:       tokensSaved,
		CompressionRatio: ratio,
		TransformsApplied: result.TransformsApplied,
	}, nil
}

// DefaultCompress is a helper function that uses a default pipeline.
func DefaultCompress(messages []models.Message, model string) (models.CompressResult, error) {
	// Build default pipeline: CacheAligner -> ContentRouter(BasicCompressor)
	pipeline := transforms.NewTransformPipeline(
		transforms.NewCacheAligner(),
		transforms.NewContentRouter(&transforms.BasicCompressor{}),
	)
	
	compressor := NewCompressor(pipeline)
	return compressor.Compress(messages, model, nil)
}
