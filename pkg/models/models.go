package models

// Message represents a single message in an LLM conversation.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"` // Simplified for MVP; real content can be a list of parts.
}

// CompressConfig controls the compression behavior.
type CompressConfig struct {
	CompressUserMessages   bool    `json:"compress_user_messages"`
	CompressSystemMessages bool    `json:"compress_system_messages"`
	ProtectRecent          int     `json:"protect_recent"`
	ProtectAnalysisContext bool    `json:"protect_analysis_context"`
	TargetRatio            *float64 `json:"target_ratio,omitempty"`
	MinTokensToCompress    int     `json:"min_tokens_to_compress"`
	KompressModel          *string `json:"kompress_model,omitempty"`
}

// NewDefaultCompressConfig returns a config with sensible defaults mirroring the Python SDK.
func NewDefaultCompressConfig() CompressConfig {
	return CompressConfig{
		CompressUserMessages:   false,
		CompressSystemMessages: true,
		ProtectRecent:          4,
		ProtectAnalysisContext: true,
		MinTokensToCompress:    250,
	}
}

// TransformResult holds the outcome of a single transform application.
type TransformResult struct {
	Messages          []Message         `json:"messages"`
	TokensBefore      int               `json:"tokens_before"`
	TokensAfter       int               `json:"tokens_after"`
	TransformsApplied []string          `json:"transforms_applied"`
	MarkersInserted   []string          `json:"markers_inserted"`
	Warnings          []string          `json:"warnings"`
	Timing            map[string]float64 `json:"timing"` // name -> duration_ms
}

// CompressResult holds the final outcome of the compression pipeline.
type CompressResult struct {
	Messages         []Message `json:"messages"`
	TokensBefore     int       `json:"tokens_before"`
	TokensAfter      int       `json:"tokens_after"`
	TokensSaved      int       `json:"tokens_saved"`
	CompressionRatio float64   `json:"compression_ratio"`
	TransformsApplied []string  `json:"transforms_applied"`
}
