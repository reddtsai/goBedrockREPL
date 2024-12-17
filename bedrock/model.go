package bedrock

type InvokeModelRequest struct {
	AnthropicVersion string    `json:"anthropic_version"`
	MaxTokens        int       `json:"max_tokens"`
	Messages         []Message `json:"messages"`
	Temperature      float64   `json:"temperature,omitempty"`
	TopP             float64   `json:"top_p,omitempty"`
	TopK             int       `json:"top_k,omitempty"`
	StopSequences    []string  `json:"stop_sequences,omitempty"`
	SystemPrompt     string    `json:"system,omitempty"`
}

type InvokeModelResponse struct {
	ID              string    `json:"id,omitempty"`
	Model           string    `json:"model,omitempty"`
	Type            string    `json:"type,omitempty"`
	Role            string    `json:"role,omitempty"`
	ResponseContent []Content `json:"content,omitempty"`
	StopReason      string    `json:"stop_reason,omitempty"`
	StopSequence    string    `json:"stop_sequence,omitempty"`
	Usage           Usage     `json:"usage,omitempty"`
}

type Message struct {
	Role    string    `json:"role,omitempty"`
	Content []Content `json:"content,omitempty"`
}

type Content struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type Usage struct {
	InputTokens  int `json:"input_tokens,omitempty"`
	OutputTokens int `json:"output_tokens,omitempty"`
}
