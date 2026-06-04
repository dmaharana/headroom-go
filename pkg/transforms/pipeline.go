package transforms

import (
	"fmt"
	"time"

	"github.com/headroom-ai/headroom-go/pkg/models"
	"github.com/headroom-ai/headroom-go/pkg/tokenizers"
)

// TransformPipeline orchestrates the application of multiple transforms.
type TransformPipeline struct {
	transforms []Transform
}

// NewTransformPipeline creates a new pipeline with the given transforms.
func NewTransformPipeline(transforms ...Transform) *TransformPipeline {
	return &TransformPipeline{
		transforms: transforms,
	}
}

// Apply executes all transforms in the pipeline.
func (p *TransformPipeline) Apply(messages []models.Message, tokenizer tokenizers.Tokenizer, config models.CompressConfig) (models.TransformResult, error) {
	currentMessages := messages
	allTransformsApplied := []string{}
	allMarkersInserted := []string{}
	allWarnings := []string{}
	allTiming := make(map[string]float64)

	tokensBeforeTotal := tokenizer.CountMessages(messages)

	pipelineStart := time.Now()

	for _, transform := range p.transforms {
		if !transform.ShouldApply(currentMessages, tokenizer, config) {
			continue
		}

		t0 := time.Now()
		result, err := transform.Apply(currentMessages, tokenizer, config)
		duration := time.Since(t0).Seconds() * 1000

		if err != nil {
			allWarnings = append(allWarnings, fmt.Sprintf("Transform %s failed: %v", transform.Name(), err))
			continue
		}

		currentMessages = result.Messages
		allTransformsApplied = append(allTransformsApplied, result.TransformsApplied...)
		allMarkersInserted = append(allMarkersInserted, result.MarkersInserted...)
		allWarnings = append(allWarnings, result.Warnings...)
		allTiming[transform.Name()] = duration
		for k, v := range result.Timing {
			allTiming[k] = v
		}
	}

	allTiming["pipeline_total"] = time.Since(pipelineStart).Seconds() * 1000
	tokensAfterTotal := tokenizer.CountMessages(currentMessages)

	return models.TransformResult{
		Messages:          currentMessages,
		TokensBefore:      tokensBeforeTotal,
		TokensAfter:       tokensAfterTotal,
		TransformsApplied: allTransformsApplied,
		MarkersInserted:   allMarkersInserted,
		Warnings:          allWarnings,
		Timing:            allTiming,
	}, nil
}
